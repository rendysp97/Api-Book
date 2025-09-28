-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  modified_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(256),
  modified_by TIMESTAMP DEFAULT NOW(),
);

-- +migrate Down
DROP TABLE users;
