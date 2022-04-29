package dto

type TransactionRequest struct {
	Amount          float64
	TransactionType string `json:"transaction_type"`
}
