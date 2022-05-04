package dto

import (
	"strings"

	"github.com/nickypangers/banking/errs"
)

const (
	WITHDRAW string = "withdraw"
	DEPOSIT  string = "deposit"
)

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"customer_id"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return strings.ToLower(r.TransactionType) == WITHDRAW
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return strings.ToLower(r.TransactionType) == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be negative")
	}

	if !r.IsTransactionTypeDeposit() && !r.IsTransactionTypeWithdrawal() {
		return errs.NewValidationError("Transaction type should be deposit or withdraw")
	}

	return nil
}
