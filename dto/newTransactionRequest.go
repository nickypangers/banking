package dto

type NewTransactionRequest struct {
	AccountId  string  `json:"account_id"`
	CustomerId string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
}
