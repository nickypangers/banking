package domain

import "github.com/nickypangers/banking/errs"

type Transaction struct {
	TransactionId   string `db:"transaction_id"`
	AccountId       string `db:"account_id"`
	Amount          float64
	TransactionType string `db:"transaction_type"`
	TransactionDate string `db:"transaction_date"`
}

type TransactionRepository interface {
	Deposit(Transaction) (*Transaction, *errs.AppError)
}
