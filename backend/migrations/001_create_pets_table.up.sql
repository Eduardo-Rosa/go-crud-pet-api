CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    species VARCHAR(100) NOT NULL,
    breed VARCHAR(100),
    age INT,
    birth_date VARCHAR(30),
    owner_name VARCHAR(100) NOT NULL
);