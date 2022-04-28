package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nickypangers/banking/errs"
	"github.com/nickypangers/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error

	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select * from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}

		logger.Error("Error while scanning customer " + err.Error())
		return nil, errs.NewUnexpectedNotFoundError("Unexpected server error")
	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:password@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
