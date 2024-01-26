package usecase

import (
	"fmt"
	"test-mnc/entity"
	"test-mnc/repository"
)

type TransactionsUsecase interface {
	RequestNewPayment(payload entity.Transaction) (entity.Transaction, error)
}

type transactionsUsecase struct {
	repo repository.TransactionsRepository
}

func NewTransactionsUsecase(repo repository.TransactionsRepository) TransactionsUsecase {
	return &transactionsUsecase{repo: repo}
}

func (t *transactionsUsecase) RequestNewPayment(payload entity.Transaction) (entity.Transaction, error) {
	transactions, err := t.repo.Create(payload)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("oppps, failed to save data transations :%v", err.Error())
	}
		return transactions, nil
}