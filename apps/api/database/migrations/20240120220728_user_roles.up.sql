CREATE TABLE user_permissions (
  	id int pk increments
    permission_id int fk permissions.id
    role_id int fk roles.id
	created_at timestamp
	updated_at timestamp
);

