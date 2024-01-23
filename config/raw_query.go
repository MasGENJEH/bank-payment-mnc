package config

const (
	InsertTransaction = `INSERT INTO transactions (customer_from, customer_to, date, amount, transaction_type, balance, description, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, date, created_at, updated_at`
)