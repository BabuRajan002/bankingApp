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
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
