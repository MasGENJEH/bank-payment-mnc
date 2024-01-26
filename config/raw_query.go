package config

const (
	InsertTransaction = `INSERT INTO transactions (customer_id, merchant_id, date, amount, transaction_type, current_balance, description, created_at, updated_at) VALUES ($1, $2, CURRENT_TIMESTAMP, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id, date, created_at, updated_at`

	SelectCustomerForLogin = `SELECT id, name, username, password FROM customers WHERE username = $1 AND password = $2`

	UpdateMerchantBalance = `UPDATE merchants SET balance = balance + $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING balance`
)