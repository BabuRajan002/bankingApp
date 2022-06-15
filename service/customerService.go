package service

import (
	"bankingApp/domain"
	"bankingApp/errs"
)

type CustomerService interface {
	GetCustomer(string) (*domain.Customer, *errs.AppError)
	GetAllCustomersByStat(string) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func (s DefaultCustomerService) GetAllCustomersByStat(stat string) ([]domain.Customer, *errs.AppError) {
	return s.repo.ByStat(stat)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
