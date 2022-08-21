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
	numDays          = 7
	numTweetsPerDay  = 10
	twitterBaseURLP1 = "https://api.twitter.com/2/tweets/search/recent?query=%22%24"
	twitterBaseURLP2 = "%22%20lang%3Aen%20-has%3Alinks&expansions=author_id&user.fields=public_metrics&tweet.fields=public_metrics&max_results="
	twitterBaseURLP3 = "&end_time="
)

var examples = []cohere.Example{
	{
		Text:  "%S is bullish keep holding until it hits",
		Label: "1",
	},
	{
		Text:  "Sell %s keeps permanent downtrend",
		Label: "0",
	},
	{
		Text:  "uptrend for %s over 150% today hold until 175%",
		Label: "1",
	},
	{
		Text:  "SHARES FALL 22% FROM JANUARY PEAK AMID TECH SELLOFF ",
		Label: "0",
	},
	{
		Text:  "%s craters, management tells u biz sucks, all sorts of red flags, and some people call that bullish?  bottom? delusionnal.. #youondrugs?",
		Label: "0",
	},
	{
		Text:  "%s Call volume spike ahead of earnings  ",
		Label: "1",
	},
	{
		Text:  "%s moving above 525 pivot today is bullish for tomorrow.",
		Label: "1",
	},
	{
		Text:  "Trend in outsourcing...workforce...is extremely strong, especially in the medical industry... %s has a lot of room to grow.",
		Label: "1",
	},
	{
		Text:  "%s in a clear downtrend with no sign of relief  ",
		Label: "0",
	},
	{
		Text:  "%s needs to hold here 500 or could see lower prices. Earnings within couple wks  		",
		Label: "0",
	},
	{
		Text:  "%s under 500 is a great long term buy.  Ignore the noise.",
		Label: "1",
	},
	{
		Text:  "The Dow industrials just wrapped up their sixth worst quarter on record. The second quarter didn't start much better",
		Label: "0",
	},
	{
		Text:  "%s - We are in at 2.55 per share.  This looks like a good one, its moving up already - we will sell at a 7% gain #stocks #insidertrading",
		Label: "1",
	},
	{
		Text:  "Sorry guys, but %s bulls need to hit the pause button.",
		Label: "0",
	},
	{
		Text:  "%s Ask is being pummeled...managers cannot miss this and afraid they are missing the train...",
		Label: "0",
	},
	{
		Text:  "%s over 23.40 here looks good to go",
		Label: "1",
	},
	{
		Text:  "%s Going on my confirmed theory selling came from those with inside info last 45 days, cat is now out of bag, no suckers left, bullish",
		Label: "1",
	},
	{
		Text:  "Big win for brokers... New Mini Options Bring Big Changes For Traders  via user %s",
		Label: "1",
	},
	{
		Text:  "%s - looks like a bear flag, we are short, will respect the flag and exit on a close above   ",
		Label: "0",
	},
	{
		Text:  "Why can't they take %s out of the index? It's holding everyone back! Slacker !!",
		Label: "0",
	},
	{
		Text:  "his approach to valuing %s extremely disciplined, well founded, well researched. At $1,600 price target, I don’t question why he’s not higher. It makes me question those that are well below his. What are they looking at?? Trading is much more of a mental game above anything. You can have the skills to read charts but how you stay mentally focused, consistent, and repetitive is how you take the next step.",
		Label: "1",
	},
	{
		Text:  "Lost 23,000$ on %s calls this wee.. oh well ???????",
		Label: "0",
	},
	{
		Text:  "If investing only in %s is too risky than why is everyone who only invested doing extremely well? Lol have you met anyone who just invested the stock ; not doing well financially? ??",
		Label: "1",
	},
	{
		Text:  "%s should be a high percent of anyones portfolio",
		Label: "1",
	},
	{
		Text:  "%s beautifully bullish",
		Label: "1",
	},
	{
		Text:  "Think I have a problem. Bought more %s at $82.55.",
		Label: "1",
	},
	{
		Text:  "%s 166 Million share float and stock is moving up on 100 share buy orders premarket i dont trust it",
		Label: "0",
	},
	{
		Text:  "Our high beta growth screen beet red. %s big negative reversal as it ran right into declining 50 day average. RETAIL getting smoked again. It's early #easymoney #markets",
		Label: "0",
	},
	{
		Text:  "If such 2 minutes of analysis is showing that spam accounts are more than 20% then why did you choose a difficult option of seeing SEC filings of Twitter? Fire your IB and DD teams. Or was 28% fall in %s share price in a month the real reason for cold feet?",
		Label: "0",
	},
	{
		Text:  "%s is temporarily halting production of its products. The stock is down ~25% on the news",
		Label: "0",
	},
	{
		Text:  "%s no longer trending finally ppl get the hint. The growth and innovation story is over... PT 375 fair value",
		Label: "0",
	},
	{
		Text:  "Nailed %s today.",
		Label: "0",
	},
	{
		Text:  "%s Cracked !!!",
		Label: "1",
	},
	{
		Text:  "%s looking for entry here.",
		Label: "1",
	},
	{
		Text:  "VP Drummond of %s sells 190 shares. He can be an outstanding trader based on his timing.",
		Label: "1",
	},
	{
		Text:  "%s don't touch it unless we get a close above 465 with volume (Earnings gap).",
		Label: "0",
	},
	{
		Text:  "%s looking good here for a long, yesterday's close as stop",
		Label: "1",
	},
	{
		Text:  "%s lookin good",
		Label: "1",
	},
	{
		Text:  "guess what happens to those chasing %s...",
		Label: "0",
	},
	{
		Text:  "%s this action in this super strong tape = Not bullish...",
		Label: "0",
	},
	{
		Text:  "%s only 10 points away from new hold time highs. This market leader continues to lead  ",
		Label: "1",
	},
	{
		Text:  "The %s drop will stop only when Wall Street and sellers stop fearing the margin drop...take a look  ",
		Label: "1",
	},
	{
		Text:  "wow %s is underrated and underloved....looking at stock prices I am sick that I missed this hate trade in 2009 .. CBS DIS",
		Label: "1",
	},
	{
		Text:  "%s of new buys this was the largest % gainer today - also was a fav as it was a pullback reversal  ",
		Label: "1",
	},
	{
		Text:  "%s is breaking out !!  ",
		Label: "1",
	},
	{
		Text:  "%s Increased Debt Causes Concern",
		Label: "0",
	},
	{
		Text:  "user: %s stop kidding yourself . there is no support around here",
		Label: "0",
	},
	{
		Text:  "iding %s down, lots to go",
		Label: "0",
	},
	{
		Text:  "seller stepped down to 89.25 in %s.",
		Label: "0",
	},
	{
		Text:  "%s decrease volume now and a red bar will send this back to 57 - 2 points to be made more if goes higher - patience and wait",
		Label: "0",
	},
}

type TwitterResult struct {
	Tweets     []Tweet `json:"data"`
	TweetUsers Users   `json:"includes"`
}

type Tweet struct {
	AuthorID     string      `json:"author_id"`
	ID           string      `json:"id"`
	Text         string      `json:"text"`
	TweetMetrics TweetMetric `json:"public_metrics"`
}

type TweetMetric struct {
	RetweetCount string `json:"retweet_count"`
	ReplyCount   string `json:"reply_count"`
	LikeCount    string `json:"like_count"`
	QuoteCount   string `json:"quote_count"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID          string     `json:"id"`
	UserMetrics UserMetric `json:"public_metrics"`
}

type UserMetric struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
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
	return twitterResponse
}

func classify(co *cohere.Client, ticker string, date time.Time, tweets []string) *cohere.ClassifyResponse {
	response, err := co.Classify(cohere.ClassifyOptions{
		Model:           config.COHERE_MODEL_ID,
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
	stockSentiments.SentimentList = make([]*pb.StockSentiment, numDays)

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
				var userMetric UserMetric
				for _, tweet := range twitterResponse.Tweets {
					// Filter out users with low follower count or following count
					for _, user := range twitterResponse.TweetUsers.Users {
						if user.ID == tweet.AuthorID {
							userMetric = user.UserMetrics
							if userMetric.FollowersCount > 100 && userMetric.FollowingCount > 20 {
								tweets = append(tweets, tweet.Text)
								break
							}
						}
					}
				}

				// classify sentiment
				co, err := cohere.CreateClient(config.COHERE_API_KEY)
				if err != nil {
					fmt.Println(err)
				}
				sentiment := 0.0
				var cohereResponse *cohere.ClassifyResponse
				if len(tweets) > 0 {
					cohereResponse = classify(co, ticker, date, tweets)
					for _, classification := range cohereResponse.Classifications {
						if i%10 == 0 {
							fmt.Println(classification.Input, classification.Prediction)
						}
						if classification.Prediction == "1" {
							sentiment += 1
						}
					}
					sentiment = sentiment * 100 / float64(len(cohereResponse.Classifications))

					// Return the first classified tweet and it's classification
					var tweetExample string
					for _, tweet := range twitterResponse.Tweets {
						if cohereResponse.Classifications[0].Input == tweet.Text {
							tweetExample = tweet.ID
							break
						}
					}
					var classificationExample string
					if cohereResponse.Classifications[0].Prediction == "1" {
						classificationExample = "Positive"
					} else {
						classificationExample = "Negative"
					}
					stock = &models.Stock{
						Ticker:                ticker,
						Date:                  date,
						Sentiment:             sentiment,
						TweetExample:          tweetExample,
						ClassificationExample: classificationExample,
					}

					// save to database
					_ = stock.CreateStock()
				}
			}

			//Append stock to result
			stockSentiments.SentimentList[numDays-(i+1)] = &pb.StockSentiment{
				Name:                  (*stock).Ticker,
				Date:                  timestamppb.New((*stock).Date),
				Sentiment:             int32((*stock).Sentiment),
				TweetExample:          (*stock).TweetExample,
				ClassificationExample: (*stock).ClassificationExample,
			}
		}(i)
	}
	workersWG.Wait()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/text")
	res, err := proto.Marshal(stockSentiments)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	w.Write(res)
}
