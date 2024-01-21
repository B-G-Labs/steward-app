CREATE TABLE IF NOT EXISTS permissions(
    id serial PRIMARY KEY NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    display_name varchar(255) NOT NULL,
    description varchar(255),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
);

