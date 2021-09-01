package main

import (
	"auctioner/apihandler"
	"auctioner/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/subosito/gotenv"
)

func main() {
	// Creating a new router
	response.Initialize()
	if gotenv.Load(".env") != nil {
		// log.Fatal("[ERROR] Failed to load the dev env file")
		fmt.Print("loaded evn")
	}
	router := mux.NewRouter()
	// Handling cors access
	router.Use(
		cors.AllowAll().Handler,
	)

	router.HandleFunc("/hello", apihandler.HelloworldDemo)
	router.HandleFunc("/registerBidder", apihandler.RegisterBidderHandler)
	router.HandleFunc("/bidderList", apihandler.BidderListHandler)
	router.HandleFunc("/bid", apihandler.BidHandler)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
