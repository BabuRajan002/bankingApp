package domain

import "bankingApp/errs"

type Customer struct {
	Id         string `db:"customer_id"`
	Name       string
	City       string
	Zipcode    string
	DateofBith string `db:"date_of_birth"`
	Status     string
}

type CustomerRepository interface {
	ById(string) (*Customer, *errs.AppError)
	ByStat(string) ([]Customer, *errs.AppError)
}
