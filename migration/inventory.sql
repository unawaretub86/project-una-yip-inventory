CREATE DATABASE inventory-db;

CREATE TABLE tech_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category VARCHAR(150),
    description TEXT,
    quantity INTEGER,
    price NUMERIC(10, 2)
);
