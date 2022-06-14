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
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
