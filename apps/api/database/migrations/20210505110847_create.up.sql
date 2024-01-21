CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  name varchar(255) UNIQUE,
  password varchar(255),
  active boolean NOT NULL DEFAULT TRUE,
  created_at timestamp DEFAULT NOW(),
  updated_at timestamp DEFAULT NOW(),
);

