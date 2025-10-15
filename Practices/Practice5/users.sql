CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100) UNIQUE,
    balance NUMERIC(10, 2)
    );

INSERT INTO users (name, email, balance) VALUES
    ('Alice', 'alice@example.com', 500.00),
    ('Bob', 'bob@example.com', 300.00);
