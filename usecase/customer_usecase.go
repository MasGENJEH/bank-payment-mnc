package usecase

import (
	"errors"
	"test-mnc/entity"
	"test-mnc/repository"
)

type CustomersUseCase interface {
	FindCustomerForLogin(username, password string) (entity.Customers, error)
}

type customersUseCase struct {
	repo repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomersUseCase {
	return &customersUseCase{repo: repo}
}

func (e *customersUseCase)FindCustomerForLogin(username, password string) (entity.Customers, error){
	if username == "" {
		return entity.Customers{}, errors.New("username harus diisi")
	}
	return e.repo.GetCustomersByUsernameForLogin(username, password)
}


