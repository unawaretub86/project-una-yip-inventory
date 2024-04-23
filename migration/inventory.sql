CREATE DATABASE inventory-db;

CREATE TABLE tech_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    category VARCHAR(150),
    description TEXT,
    quantity INTEGER,
    price NUMERIC(10, 2)
);

ALTER TABLE tech_items ADD CONSTRAINT unique_name_constraint UNIQUE name;
ALTER TABLE tech_items ALTER COLUMN category SET NOT NULL;
ALTER TABLE tech_items ALTER COLUMN quantity SET NOT NULL;
ALTER TABLE tech_items ALTER COLUMN description SET NOT NULL;
ALTER TABLE tech_items ALTER COLUMN price SET NOT NULL;
