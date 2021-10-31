-- Tables creation
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    address TEXT NOT NULL,
    password TEXT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
CREATE TABLE IF NOT EXISTS items (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    quantity INTEGER DEFAULT 0,
    price INTEGER NOT NULL
);

-- Seeding items
INSERT INTO items(name, quantity, price)
VALUES ('Iphone 12 Pro', 2, 12500000);
INSERT INTO items(name, quantity, price)
VALUES ('Samsung Max Pro', 1, 2100000);
INSERT INTO items(name, quantity, price)
VALUES ('Screen Guard 5.5"', 1, 50000);
