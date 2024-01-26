CREATE DATABASE bank_payment_db;

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    birth_date DATE NOT NULL,
    address TEXT NOT NULL
)

CREATE TABLE merchants (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  merchantCode VARCHAR(255) NOT NULL,
  balance DOUBLE PRECISION NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE SEQUENCE transactions_id_seq START 1;

CREATE TABLE transactions (
    id VARCHAR(10) PRIMARY KEY DEFAULT 'TRX-' || LPAD(NEXTVAL('transactions_id_seq')::TEXT, 5, '0'),
    customer_id INT NOT NULL,
    merchant_id INT NOT NULL,
    date DATE NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    transaction_type VARCHAR(10),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id),
    FOREIGN KEY (merchant_id) REFERENCES merchants(id)
)