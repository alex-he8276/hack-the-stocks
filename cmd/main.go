package main

import (
	"log"
	"net/http"

	"github.com/alex-he8276/hack-the-stocks/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal(err)
	}
}
