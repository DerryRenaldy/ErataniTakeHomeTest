-- migrate:up

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    country VARCHAR(100),
    credit_card_type VARCHAR(50),
    credit_card_number VARCHAR(20),
    first_name VARCHAR(100),
    last_name VARCHAR(150)
);

-- Create user_transactions table
CREATE TABLE IF NOT EXISTS user_transactions (
    id BIGSERIAL PRIMARY KEY,
    id_user INT NOT NULL,
    total_buy BIGINT NOT NULL
);

-- Add foreign key constraint
ALTER TABLE user_transactions
DROP CONSTRAINT IF EXISTS fk_user_transactions_user;

ALTER TABLE user_transactions
ADD CONSTRAINT fk_user_transactions_user
FOREIGN KEY (id_user)
REFERENCES users(id);

-- migrate:down

DROP TABLE IF EXISTS user_transactions;
DROP TABLE IF EXISTS users;
