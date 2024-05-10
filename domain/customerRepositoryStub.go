package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Sajjad", "Charsadda", "24420", "23-03-1993", "1"},
		{"1002", "Abrar", "NY", "23", "23-03-03", "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
