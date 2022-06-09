package domain

type Customer struct {
	Id         string
	Name       string
	City       string
	Zipcode    string
	DateofBith string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
