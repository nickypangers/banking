package domain

import "github.com/nickypangers/banking/errs"

type Customer struct {
	Id          string `db:"customer_id" json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	DateofBirth string `db:"date_of_birth" json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
