package domain

import (
	"github.com/nickypangers/banking/dto"
	"github.com/nickypangers/banking/errs"
)

type Transaction struct {
	TransactionId   string `db:"transaction_id"`
	AccountId       string `db:"account_id"`
	Amount          float64
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		TransactionDate: t.TransactionDate,
	}
}
