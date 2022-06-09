package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil

}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Babu", City: "Trichy", Zipcode: "621312", DateofBith: "1991-06-02"},
		{Id: "1002", Name: "Vinna", City: "Dindigul", Zipcode: "600100", DateofBith: "1993-05-29"},
	}
	return CustomerRepositoryStub{customers}
}
