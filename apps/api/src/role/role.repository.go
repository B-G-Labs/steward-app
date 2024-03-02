package role

import (
	"context"

	"github.com/uptrace/bun"
)

type RoleRepository struct {
	db *bun.DB
	ctx context.Context
}

// ???
// type repository interface {
// 	ListRoles() ([]Role, error)
// 	DeleteRoles(role Role) (int64, error)
// 	CreateRole(role Role) (int64, error)
// 	UpdateRole(role Role) (Role, error)
// }

func NewRoleRepository(database *bun.DB, ctx context.Context) RoleRepository {
	return RoleRepository{
		db: database,
		ctx: ctx,
	}
}

func (r *RoleRepository) getRoleById(id int64) (Role, error) {
	var role Role
	
	err := r.db.NewSelect().
		Model(&role).
		Where("id = ?", id).
		Scan(r.ctx)
	
	if err != nil {
		return Role{}, err
	}

	return role, nil
}

func (r *RoleRepository) CreateRole(role Role) (int64, error) {
	model :=  &Role{
		Name: role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
	}
	
	err := r.db.NewInsert().
		Model(model).
		Returning("id").
		Scan(r.ctx)
	
	if err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (r *RoleRepository) DeleteRole(role Role) (int64, error) {
	result, err := r.db.NewDelete().
		Model(&role).
		Where("id = ?", role.ID).
		Exec(r.ctx)
	
	if err != nil {
		return 0, err
	}
	
	rowsAffected, err := result.RowsAffected()
	
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *RoleRepository) ListRoles() ([]Role, error) {
	var roles []Role
	
	err := r.db.NewSelect().
		Model(&roles).
		Scan(r.ctx)
	
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepository) UpdateRole(role Role) (Role, error) {
	model :=  &Role{
		ID: role.ID,
		Name: role.Name,
		DisplayName: role.DisplayName,
		Description: role.Description,
	}
	
	_, err := r.db.NewUpdate().
		Model(model).
		Where("id = ?", role.ID).
		Exec(r.ctx)
	
	if err != nil {
		return Role{}, err
	}

	return role, nil
}