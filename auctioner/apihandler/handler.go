package apihandler

import (
	"auctioner/models"
	"auctioner/response"
	"auctioner/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

var count = 0

var bidderList = make([]models.RegisterBidder, 0)

func HelloworldDemo(w http.ResponseWriter, r *http.Request) {
	obj := make(map[string]interface{})
	obj["message"] = "Hello world"

	count += 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprint("Hello world ", count),
	})
}

func RegisterBidderHandler(w http.ResponseWriter, r *http.Request) {

	var registerRequestData models.RegisterBidder
	// Decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&registerRequestData); err != nil {
		// Failed to decode the request body
		response.BadRequest(w, "110")
	} else {

		bidderList = append(bidderList, registerRequestData)

		response.Success(w, "111", []interface{}{})
	}
}

func BidderListHandler(w http.ResponseWriter, r *http.Request) {

	response.Success(w, "111", bidderList)
}

func BidHandler(w http.ResponseWriter, r *http.Request) {

	var bidRequestData models.RequestAuctioner
	// Decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&bidRequestData); err != nil {
		// Failed to decode the request body
		response.BadRequest(w, "110")
	} else {

		maxBidder, _ := utils.GetMaxValueBidder(bidderList)

		response.Success(w, "111", maxBidder)
	}
}
