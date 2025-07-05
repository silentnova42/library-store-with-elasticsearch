CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    book_name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    number_of_copies INT NOT NULL,
    author VARCHAR(255) NOT NULL
);