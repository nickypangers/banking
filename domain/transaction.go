package domain

import "github.com/nickypangers/banking/errs"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	Deposit(Transaction) (*Transaction, *errs.AppError)
}
