CREATE TABLE customers(
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100),
    phone VARCHAR(10)
);