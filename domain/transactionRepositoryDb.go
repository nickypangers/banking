package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/logger"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert transaction id " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	t.TransactionId = strconv.FormatInt(id, 10)
	return &t, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}
