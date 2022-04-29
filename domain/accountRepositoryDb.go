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

func (d AccountRepositoryDb) ById(customerId, accountId string) (*Account, *errs.AppError) {
	accountSql := "select * from accounts where customer_id = ? AND account_id = ?"

	var a Account
	err := d.client.Get(&a, accountSql, customerId, accountId)
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

func (d AccountRepositoryDb) UpdateAmount(a Account) (*Account, *errs.AppError) {

	sqlUpdate := "UPDATE accounts SET amount = ? WHERE customer_id = ? AND account_id = ?"

	result, err := d.client.Exec(sqlUpdate, a.Amount, a.CustomerId, a.AccountId)
	if err != nil {
		logger.Error("Error while updating account " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("Error while updating account " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}

	if rowsAffected == 0 {
		logger.Error("Error while updating account " + err.Error())
		return nil, errs.NewNotFoundError("Account not found")
	}

	newA, appError := d.ById(a.CustomerId, a.AccountId)
	if appError != nil {
		return nil, appError
	}

	return newA, nil

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

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
