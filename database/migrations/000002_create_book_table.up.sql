-- +migrate Up
CREATE TABLE book (
  id SERIAL PRIMARY KEY, 
  title VARCHAR(100) NOT NULL,
  description VARCHAR(100) NOT NULL,
  image_url VARCHAR(100) NOT NULL,
  release_year INTEGER,
  price INTEGER,
  total_page INTEGER,
  thickness VARCHAR(50),
  category_id INTEGER NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  modified_at TIMESTAMP DEFAULT NOW(),
  created_by VARCHAR(256),
  modified_by TIMESTAMP DEFAULT NOW(),
  CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE book;