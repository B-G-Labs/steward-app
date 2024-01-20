CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name varchar(255) UNIQUE,
  password varchar(255),
  active boolean,
	created_at timestamp,
	updated_at timestamp
);

