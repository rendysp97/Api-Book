-- +migrate Up

CREATE TABLE categories (
id SERIAL PRIMARY KEY,
 name VARCHAR(100) NOT NULL,
 created_at TIMESTAMP DEFAULT NOW(),
 created_by VARCHAR(256),
 modified_at TIMESTAMP DEFAULT NOW(),
 modified_by TIMESTAMP DEFAULT NOW()

);

-- +migrate Down
DROP TABLE categories;
