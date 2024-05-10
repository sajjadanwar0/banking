package domain

import (
	"banking/dto"
	"banking/errs"
)

type Customer struct {
	Id          string `json:"Id" db:"customer_id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "Active"
	if c.Status == "0" {
		statusAsText = "Inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
