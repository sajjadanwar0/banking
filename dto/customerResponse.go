package dto

type CustomerResponse struct {
	Id          string `json:"customer_id" db:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode" db:"zipcode"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth"`
	Status      string `json:"status"`
}
