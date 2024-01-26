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
	var merchant entity.Merchants
	err := t.db.QueryRow(config.UpdateMerchantBalance, 
		payload.Amount,
		payload.MerchantId).Scan(&merchant.Balance)
	
	if err != nil {
		log.Println("updateMerchantBalanceRepository.QueryRow: ", err.Error())
		return entity.Transaction{}, err
	}
		
	err = t.db.QueryRow(config.InsertTransaction,
		payload.CustomerId,
		payload.MerchantId,
		payload.Amount,
		payload.TransactionType,
		merchant.Balance,
		payload.Description).Scan(&transaction.Id, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)

	if err != nil {
		log.Println("transactionRepository.QueryRow: ", err.Error())
		return entity.Transaction{}, err
	}


	transaction.CustomerId = payload.CustomerId
	transaction.MerchantId = payload.MerchantId
	transaction.Amount = payload.Amount
	transaction.TransactionType = payload.TransactionType
	transaction.Balance = merchant.Balance
	transaction.Description = payload.Description

	return transaction, nil	
}

