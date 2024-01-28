create table role_permissions (
    id serial primary key not null,
    permission_id int not null,
    role_id int not null,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    constraint fk_permission foreign key (permission_id) references permissions (id) on delete cascade,
    constraint fk_user foreign key (role_id) references roles (id) on delete cascade
);

