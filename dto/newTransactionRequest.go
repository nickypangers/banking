package dto

type NewTransactionRequest struct {
	AccountId  string
	CustomerId string
	Amount     float64
}
