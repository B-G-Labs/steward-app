package permission

import (
	"context"

	"github.com/uptrace/bun"
)

type PermissionRepository struct {
	db  *bun.DB
	ctx context.Context
}

type repository interface {
	ListPermissions() ([]Permission, error)
	UpdatePermission(p Permission) (Permission, error)
	CreatePermission(p Permission) (int64, error)
	GetPermissionsById(id int64) (Permission, error)
	DeletePermission(id int64) (int64, error)
	AssignPermissionToRole(roleId int64, permissionId int64) (bool, error)
}

func NewPermissionRepository(database *bun.DB, ctx context.Context) PermissionRepository {
	return PermissionRepository{
		db:  database,
		ctx: ctx,
	}
}

func (r *PermissionRepository) AssignPermissionToRole(roleId int64, permissionId int64) (bool, error) {
	values := map[string]interface{}{
		"role_id":       roleId,
		"permission_id": permissionId,
	}

	_, err := r.db.NewInsert().
		Model(&values).
		TableExpr("role_permissions").
		Exec(r.ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *PermissionRepository) ListPermissions() ([]Permission, error) {
	var permissions []Permission

	err := r.db.NewSelect().
		Model(&permissions).
		Scan(r.ctx)

	if err != nil {
		return permissions, err
	}

	return permissions, nil
}

func (r *PermissionRepository) GetPermissionsById(id int64) (Permission, error) {
	permission := new(Permission)

	err := r.db.NewSelect().
		Model(permission).
		Where("id = ?", id).
		Scan(r.ctx)

	if err != nil {
		return *permission, err
	}

	return *permission, nil
}

func (r *PermissionRepository) CreatePermission(p Permission) (int64, error) {
	model := &Permission{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}

	result, err := r.db.NewInsert().
		Model(model).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PermissionRepository) UpdatePermission(p Permission) (Permission, error) {
	inputPermission := &Permission{
		ID:          p.ID,
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
	}

	_, err := r.db.NewUpdate().
		Model(inputPermission).
		WherePK().
		Returning("*").
		Exec(r.ctx)

	if err != nil {
		return *inputPermission, err
	}

	return *inputPermission, nil
}

func (r *PermissionRepository) DeletePermission(id int64) (int64, error) {
	permission := new(Permission)

	result, err := r.db.NewDelete().
		Model(permission).
		Where("id = ?", id).
		Exec(r.ctx)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
