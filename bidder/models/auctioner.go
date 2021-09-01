package models

type ResponseBidder struct {
	BidderID string `json:"bidder_id"`
	BidValue int    `json:"bid_value"`
}

type RegisterBidder struct {
	Url string `json:"url"    validate:"required"`
}
