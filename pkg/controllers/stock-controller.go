package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/cohere-ai/cohere-go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/alex-he8276/hack-the-stocks/pkg/config"
	"github.com/alex-he8276/hack-the-stocks/pkg/models"
	pb "github.com/alex-he8276/hack-the-stocks/pkg/proto"
	"github.com/alex-he8276/hack-the-stocks/pkg/utils"
)

const (
	numDays          = 3
	numTweetsPerDay  = 10
	twitterBaseURLP1 = "https://api.twitter.com/2/tweets/search/recent?query=%22"

	twitterBaseURLP2 = "%22%20lang%3Aen%20-has%3Alinks&expansions=author_id&user.fields=public_metrics&tweet.fields=public_metrics&max_results="
	twitterBaseURLP3 = "&end_time="
)

var examples = []cohere.Example{
	{
		Text:  "The stock market is going up",
		Label: "positive",
	},
	{
		Text:  "Sell or you'll regret it",
		Label: "negative",
	},
	{
		Text:  "The stock market is going down",
		Label: "negative",
	},
	{
		Text:  "@elonmusk I feel so bad for both $TSLA and $TSLAQ. You all live the most pathetic lives possible. 24/7 documenting crap that never adds up to anything. It’s turned into conspiracy theories, fuel by confirmation bias, disdain, and vitriol. Masacadaded by avatars and themed handles.",
		Label: "negative",
	},
	{
		Text:  "What does this company do?",
		Label: "neutral",
	},
	{
		Text:  "Tesla should be a high percent of anyone’s portfolio #TSLA",
		Label: "positive",
	},
	{
		Text:  "The potential for TSLA is huge",
		Label: "positive",
	},
	{
		Text:  "The CEO is a complete moron, stay away if you value your wealth",
		Label: "negative",
	},
	{
		Text:  "This stock is a great buy, returned me over 80% in the past 6 years",
		Label: "positive",
	},
	{
		Text:  "DIAMOND HANDS",
		Label: "positive",
	},
	{
		Text:  "Company reported 0% change in revenue",
		Label: "neutral",
	},
	{
		Text:  "This stock is going bust, dont buy",
		Label: "negative",
	},
}

type TwitterResult struct {
	Tweets []Tweet `json:"data"`
	TweetUsers Users `json:includes`
}

type Tweet struct {
	AuthorID string `json:author_id`
	ID   string `json:"id"`
	Text string `json:"text"`
	TweetMetrics TweetMetric `json:public_metrics`
}

type TweetMetric struct {
	RetweetCount string `retweet_count`
	ReplyCount string `reply_count`
	LikeCount string `like_count`
	QuoteCount string `quote_count`
}

type Users struct {
	Users []User `json:users`
}

type User struct {
	ID   string `json:"id"`
	UserMetrics UserMetric `json:public_metrics`
}

type UserMetric struct {
	FollowersCount   int `json:"followers_count"`
	FollowingCount   int `json:"following_count"`
}

func checkDatabase(ticker string, date time.Time) (*models.Stock, bool) {
	stock := models.GetStockByTickerAndDate(ticker, date)
	if stock.Ticker == "" {
		return stock, false
	}
	return stock, true
}

func retrieveTweets(ticker string, date time.Time) TwitterResult {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
	// curl "https://api.twitter.com/2/tweets/search/recent?query=%22$TICKER%22" -H "Authorization: Bearer $BEARERTOKEN"
	url := fmt.Sprintf("%s%s%s%s%s%s", twitterBaseURLP1, ticker, twitterBaseURLP2, strconv.Itoa(numTweetsPerDay), twitterBaseURLP3, date.Format(time.RFC3339))
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.TWITTER_BEARER_TOKEN))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("failed retreiving tweets", err)
	}
	defer resp.Body.Close()
	twitterResponse := TwitterResult{}
	utils.ParseBody(resp, &twitterResponse)

	// TODO: Filter Tweet Creator by Followers
	fmt.Println(twitterResponse)
	return twitterResponse
}

func classify(ticker string, date time.Time, tweets []string) *cohere.ClassifyResponse {
	co, err := cohere.CreateClient(config.COHERE_API_KEY)
	if err != nil {
		fmt.Println(err)
	}

	response, err := co.Classify(cohere.ClassifyOptions{
		Model:           "medium",
		OutputIndicator: "Classify this tweet about a stock",
		TaskDescription: "Classify these tweets about a stock as positive, negative or neutral",
		Inputs:          tweets,
		Examples:        examples,
	})

	if err != nil {
		fmt.Println(err)
	}

	return response
}

func GetStockSentiment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticker := vars["ticker"]
	stockSentiments := &pb.ListStockSentiment{}
	

	workersWG := sync.WaitGroup{}
	for i := 0; i < numDays; i++ {
		workersWG.Add(1)

		// Worker for a given date i days in the past
		go func(i int) {
			defer workersWG.Done()

			date := time.Now().Add(-24*time.Duration(i)*time.Hour - time.Minute)

			stock, exists := checkDatabase(ticker, date)

			if !exists {
				// Retrieve tweets from twitter api
				twitterResponse := retrieveTweets(ticker, date)

				// get text from twitter response
				tweets := make([]string, 0)
				for _, tweet := range twitterResponse.Tweets {
					tweets = append(tweets, tweet.Text)
				}
				// classify sentiment
				cohereResponse := classify(ticker, date, tweets)
				sentiment := 0.0
				for _, classification := range cohereResponse.Classifications {
					fmt.Println(classification.Input, classification.Prediction)
					if classification.Prediction == "positive" {
						sentiment += 1
					}
				}
				sentiment = sentiment * 100 / float64(len(cohereResponse.Classifications))
				stock = &models.Stock{
					Ticker:    ticker,
					Date:      date,
					Sentiment: sentiment,
				}
				
				// save to database
				_ = stock.CreateStock()
			}
			//Append stock to result
			stockSentiments.SentimentList = append(stockSentiments.SentimentList, &pb.StockSentiment{
			Name:      (*stock).Ticker,
			Date:      timestamppb.New((*stock).Date),
			Sentiment: int32((*stock).Sentiment),
			})

		}(i)
	}
	workersWG.Wait()
	// w.Write([]byte(fmt.Sprintf("%s", ticker)))

	res, err := proto.Marshal(stockSentiments)
	// res, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	w.Write(res)
}
