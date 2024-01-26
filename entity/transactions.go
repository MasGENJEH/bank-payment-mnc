package entity

import "time"

type Transaction struct {
	Id              string      `json:"id"`
	CustomerId []Customers `json:"customers_id"`
	MerchantId     string `json:"merchant_id"`
	Date time.Time `json:"date"`
	Amount          float64     `json:"amount"`
	TransactionType string      `json:"transaction_type"`
	Balance         float64     `json:"balance"`
	Description     string      `json:"description"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}