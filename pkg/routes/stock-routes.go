package routes

import (
	"github.com/alex-he8276/hack-the-stocks/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/stock/{ticker}", controllers.GetStockSentiment).Methods("GET")
}
