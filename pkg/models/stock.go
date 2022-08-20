package models

import (
	"time"

	"github.com/alex-he8276/hack-the-stocks/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Stock struct {
	// gorm.Model
	Ticker                string    `gorm:""json:"ticker"`
	Date                  time.Time `json:"date"`
	Sentiment             float64   `json:"sentiment"`
	TweetExample          string    `json:"tweet_example"`
	ClassificationExample string    `json:"classification_example"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Stock{})
}

func GetAllStocks() []Stock {
	var stocks []Stock
	db.Find(&stocks)
	return stocks
}

func GetStockByTickerAndDate(ticker string, date time.Time) *Stock {
	var stock Stock
	_ = db.Where("Ticker = ? AND Date = ?", ticker, date).Find(&stock)
	return &stock
}

func (s *Stock) CreateStock() *Stock {
	db.NewRecord(s)
	db.Create(&s)
	return s
}
