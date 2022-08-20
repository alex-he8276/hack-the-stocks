package config

import (
	"fmt"

	"github.com/cohere-ai/cohere-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var co *cohere.Client

func ConnectDB() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/happy-db?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
}

func GetDB() *gorm.DB {
	return db
}

func ConnectCohere() {
	var err error
	co, err = cohere.CreateClient(COHERE_API_KEY)
	if err != nil {
		fmt.Println(err)
	}
}

func GetCohereClient() *cohere.Client {
	return co
}
