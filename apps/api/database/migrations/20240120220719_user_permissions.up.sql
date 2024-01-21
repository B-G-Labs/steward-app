CREATE TABLE IF NOT EXISTS user_permissions (
    id serial PRIMARY KEY NOT NULL,
    permission_id int NOT NULL,
    user_id int NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
    CONSTRAINT fk_permission FOREIGN KEY (permission_id) REFERENCES permissions (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

