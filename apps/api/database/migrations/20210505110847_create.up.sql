CREATE TABLE users (
  ID SERIAL PRIMARY KEY,
  Name varchar(255) UNIQUE,
  Password varchar(255),
  Active boolean,
);

