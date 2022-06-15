package domain

import "bankingApp/errs"

type Customer struct {
	Id         string
	Name       string
	City       string
	Zipcode    string
	DateofBith string
	Status     string
}

type CustomerRepository interface {
	ById(string) (*Customer, *errs.AppError)
	ByStat(string) ([]Customer, *errs.AppError)
}
