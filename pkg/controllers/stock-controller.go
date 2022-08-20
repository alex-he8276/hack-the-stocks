package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/alex-he8276/hack-the-stocks/pkg/config"
	"github.com/alex-he8276/hack-the-stocks/pkg/models"
	"github.com/alex-he8276/hack-the-stocks/pkg/utils"
	"github.com/cohere-ai/cohere-go"
	"github.com/gorilla/mux"
)

const (
	numDays          = 1
	tweets           = 10
	twitterBaseURLP1 = "https://api.twitter.com/2/tweets/search/recent?query=%22"
	twitterBaseURLP2 = "%22%20lang%3Aen%20-has%3Alinks&max_results=10&end_time="
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
}

type Tweet struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// type TwitterMetadata struct {
// 	NewestID    string `json:"newest_id"`
// 	OldestID    string `json:"oldest_id"`
// 	ResultCount int    `json:"result_count"`
// }

// type CohereResponse struct {
// 	StockClassifications StockClassification `json:"results"`
// }

// type StockClassification struct {
// 	Text       string `json:"text"`
// 	Prediction string `json:"prediction"`
// }

type StockSentimentAndPrice struct {
	Stock [numDays]models.Stock
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
	url := fmt.Sprintf("%s%s%s%s", twitterBaseURLP1, ticker, twitterBaseURLP2, date.Format(time.RFC3339))
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

func GetStockSentimentAndPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticker := vars["ticker"]
	result := StockSentimentAndPrice{}

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
				fmt.Println(sentiment, ticker)
				sentiment = sentiment * 100 / float64(len(cohereResponse.Classifications))
				fmt.Println(sentiment, ticker)
				stock = &models.Stock{
					Ticker:    ticker,
					Date:      date,
					Sentiment: sentiment,
				}

				// save to database
				_ = stock.CreateStock()
			}

			// save the result to return
			result.Stock[i] = *stock
		}(i)
	}
	workersWG.Wait()
	w.Write([]byte(fmt.Sprintf("%s", ticker)))

	// TODO: protobuf
	res, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	w.Write(res)
}
