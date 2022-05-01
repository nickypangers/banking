package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) ById(accountId string) (*Account, *errs.AppError) {
	accountSql := "select * from accounts where account_id = ?"

	var a Account
	err := d.client.Get(&a, accountSql, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning account " + err.Error())
			return nil, errs.NewNotFoundError("Account not found")
		}

		logger.Error("Error while scanning account " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	return &a, nil
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while saving account " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert account id " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction: " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert transaction id " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	account, appErr := d.ById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	t.Amount = account.Amount
	return &t, nil

}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
