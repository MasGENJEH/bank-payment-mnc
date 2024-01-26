package entity

import "time"

type Transaction struct {
	Id              string      `json:"id"`
	CustomerId string `json:"customers_id"`
	MerchantId     string `json:"merchant_id"`
	Date time.Time `json:"date"`
	Amount          float64     `json:"amount"`
	TransactionType string      `json:"transaction_type"`
	CurrentBalance         float64     `json:"current_balance"`
	Description     string      `json:"description"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}