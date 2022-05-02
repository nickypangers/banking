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
	// starting the database transaction block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected database error")
	}
	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.ById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
