CREATE TABLE user_permissions (
  	id int pk increments
    permission_id int fk permissions.id
    user_id int fk users.id
	created_at timestamp
	updated_at timestamp
);

