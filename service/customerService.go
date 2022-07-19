package service

import (
	"bankingApp/domain"
	"bankingApp/dto"
	"bankingApp/errs"
)

type CustomerService interface {
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	GetAllCustomersByStat(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) GetAllCustomersByStat(stat string) ([]domain.Customer, *errs.AppError) {
	return s.repo.ByStat(stat)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
