CREATE TABLE IF NOT EXISTS products (
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price REAL NOT NULL
);