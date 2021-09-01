package main

import (
	"bidder/apihandler"
	"bidder/response"
	"bidder/utils"
	"fmt"
	"net/http"
	"os"

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

	err1 := utils.SendRegisterRequest(os.Getenv("URL"), os.Getenv("HOST"), os.Getenv("PORT"))
	if err1 != nil {
		panic(err1)
	}

	router := mux.NewRouter()
	// Handling cors access
	router.Use(
		cors.AllowAll().Handler,
	)

	router.HandleFunc("/hello", apihandler.HelloworldDemo)
	router.HandleFunc("/bid", apihandler.BidHandler)

	err := http.ListenAndServe(":7000", router)

	if err != nil {
		panic(err)
	}
}
