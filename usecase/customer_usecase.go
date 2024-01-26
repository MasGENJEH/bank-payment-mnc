package usecase

import (
	"errors"
	"test-mnc/entity"
	"test-mnc/entity/dto"
	"test-mnc/repository"
)

type CustomersUseCase interface {
	FindCustomerForLogin(username, password string) (entity.Customers, error)
	DeleteTokenForLogout(token dto.AuthResponseDto) error
}

type customersUseCase struct {
	repo repository.CustomerRepository
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomersUseCase {
	return &customersUseCase{repo: repo}
}

func (e *customersUseCase) FindCustomerForLogin(username, password string) (entity.Customers, error){
	if username == "" {
		return entity.Customers{}, errors.New("username harus diisi")
	}
	return e.repo.GetCustomersByUsernameForLogin(username, password)
}

func (e *customersUseCase) DeleteTokenForLogout(token dto.AuthResponseDto) error {
	return e.repo.DeleteToken(token)
}


