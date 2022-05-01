package dto

import (
	"strings"

	"github.com/nickypangers/banking/errs"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"customer_id"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be negative")
	}

	if strings.ToLower(t.TransactionType) != "deposit" && strings.ToLower(t.TransactionType) != "withdraw" {
		return errs.NewValidationError("Transaction type should be deposit or withdraw")
	}

	return nil
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return strings.ToLower(t.TransactionType) == "withdraw"
}
