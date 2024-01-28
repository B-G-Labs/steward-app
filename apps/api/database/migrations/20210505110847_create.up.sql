create table if not exists users (
  id serial primary key,
  name varchar(255) unique,
  password varchar(255),
  active boolean not null default true,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);