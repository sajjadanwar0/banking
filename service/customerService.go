package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err

	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
