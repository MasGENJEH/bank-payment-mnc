package repository

import (
	"database/sql"
	"log"
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/entity/dto"
)

type CustomerRepository interface {
	GetCustomersByUsernameForLogin(username, password string) (entity.Customers, error)
	DeleteToken(user dto.AuthResponseDto) error
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (e *customerRepository) GetCustomersByUsernameForLogin(username, password string) (entity.Customers, error) {
	var customer entity.Customers
	err := e.db.QueryRow(config.SelectCustomerForLogin, username, password).Scan(
		&customer.Id,
		&customer.Name,
		&customer.Username,
		&customer.Password)
	if err != nil {
		log.Println("customerRepository.GetCustomerByID.QueryRow: ", err.Error())
		return entity.Customers{}, err
	}
	return customer, nil
}

func (c *customerRepository) DeleteToken(user dto.AuthResponseDto) error {
	user.Token = ""
	return nil
}