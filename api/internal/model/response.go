package model

type ErrorResponse struct {
	Error string
}
type SuccessResponse struct {
	Success string
}
type RateValue struct {
	Rate uint32 `json:"rate"`
}
