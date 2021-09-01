package models

type RequestAuctioner struct {
	AuctionID string `json:"auction_id"    validate:"required"`
}
type ResponseBidder struct {
	BidderID string `json:"bidder_id"`
	BidValue int    `json:"bid_value"`
}
type ResponseTemplate struct {
	ResponseCode    string         `json:"response_code"`
	ResponseMessage string         `json:"response_message"`
	Data            ResponseBidder `json:"data"`
}
type RegisterBidder struct {
	Url string `json:"url"    validate:"required"`
}
