CREATE DATABASE bank_payment_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;


CREATE TABLE customers (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    birth_date DATE NOT NULL,
    address TEXT NOT NULL,
    contact VARCHAR(12) NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
)

CREATE SEQUENCE transactions_id_seq START 1;

CREATE TABLE transactions (
    id VARCHAR(10) PRIMARY KEY DEFAULT 'TRX-' || LPAD(NEXTVAL('transactions_id_seq')::TEXT, 5, '0'),
    customer_from uuid NOT NULL,
    customer_to uuid NOT NULL,
    date DATE NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    transaction_type VARCHAR(10),
    balance DOUBLE PRECISION NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_from) REFERENCES customers(id),
    FOREIGN KEY (customer_to) REFERENCES customers(id)
)