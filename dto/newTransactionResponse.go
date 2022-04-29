package dto

type NewTransactionResponse struct {
	TransactionId   string `json:"transaction_id"`
	TransactionDate string `json:"transaction_date"`
}
