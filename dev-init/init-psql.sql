CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
