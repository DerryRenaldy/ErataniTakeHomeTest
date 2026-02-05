-- migrate:up

-- Seed users from CSV
COPY users(id, country, credit_card_type, credit_card_number, first_name, last_name) FROM '/seed-data/users.csv' WITH (FORMAT csv, HEADER true);

-- Seed user_transactions from CSV
COPY user_transactions(id, id_user, total_buy) FROM '/seed-data/user_transactions.csv' WITH (FORMAT csv, HEADER true);

-- Update serial sequences
SELECT setval(
    pg_get_serial_sequence('users', 'id'),
    COALESCE((SELECT MAX(id) FROM users), 0) + 1,
    true
);

SELECT setval(
    pg_get_serial_sequence('user_transactions', 'id'),
    COALESCE((SELECT MAX(id) FROM user_transactions), 0) + 1,
    true
);

-- migrate:down

-- No down migration needed for seed data
