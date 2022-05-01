CREATE TABLE IF NOT EXISTS products
(
    id SERIAL PRIMARY KEY,
    external_id INT,
    type VARCHAR(255),
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS products_translated
(
    id SERIAL PRIMARY KEY,
    external_id INT,
    type VARCHAR(255),
    name VARCHAR(255)
);