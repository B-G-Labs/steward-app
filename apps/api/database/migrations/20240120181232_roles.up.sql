create table roles (
  	id serial primary key not null,
	name varchar(255) unique not null,
	display_name varchar(255) not null,
	description varchar(255),
    created_at timestamp default now(),
    updated_at timestamp default now()
);

