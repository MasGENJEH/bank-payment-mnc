package repository

import (
	"database/sql"
	"log"
	"test-mnc/config"
	"test-mnc/entity"
)

type TransactionsRepository interface {
}

type transactionsRepository struct {
	db *sql.DB
}

func NewTransactionsRepository(db *sql.DB) TransactionsRepository {
	return &transactionsRepository{db: db}
}

func (t *transactionsRepository) Create(payload entity.Transaction) (entity.Transaction, error) {
	var transaction entity.Transaction

	err := t.db.QueryRow(config.InsertTransaction,
		payload.CustomersFrom,
		payload.CustomersTo,
		payload.Amount,
		payload.TransactionType,
		payload.Balance,
		payload.Description).Scan(&transaction.Id, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		log.Println("transactionRepository.QueryRow: ", err.Error())
		return entity.Transaction{}, err
	}

	transaction.CustomersFrom = payload.CustomersFrom
	transaction.CustomersTo = payload.CustomersTo
	transaction.Amount = payload.Amount
	transaction.TransactionType = payload.TransactionType
	transaction.Balance = payload.Balance
	transaction.Description = payload.Description

	return transaction, nil	
}

