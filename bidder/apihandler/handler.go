package apihandler

import (
	"bidder/models"
	"bidder/response"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HelloworldDemo(w http.ResponseWriter, r *http.Request) {
	obj := make(map[string]interface{})
	obj["message"] = "Hello world"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(struct {
		Message string `json:"message"`
	}{
		Message: "Hello world",
	})
}

func BidHandler(w http.ResponseWriter, r *http.Request) {
	delay, err := strconv.Atoi(os.Getenv("DELAY"))
	if err != nil {
		response.BadRequest(w, "110")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	time.Sleep(time.Duration(delay) * time.Millisecond)
	response.Success(w, "111", models.ResponseBidder{
		BidderID: os.Getenv("BIDDER"),
		BidValue: rand.Intn(1000),
	})
}
