package permission

import (
	"context"

	"github.com/uptrace/bun"
)

type PermissionService interface {
	ListPermissions() ([]Permission, error)
	UpdatePermission(p Permission) (Permission, error)
	CreatePermission(p Permission) (int64, error)
	GetPermissionsById(id int64) (Permission, error)
	DeletePermission(id int64) (int64, error)
	AssignPermissionToRole(roleId int64, permissionId int64) (bool, error)
}

type service struct {
	repository PermissionRepository
	ctx        context.Context
}

func NewService(database *bun.DB, ctx context.Context) PermissionService {
	userRepo := NewPermissionRepository(database, ctx)

	return &service{
		repository: userRepo,
		ctx:        ctx,
	}
}

func (s *service) AssignPermissionToRole(roleId int64, permissionId int64) (bool, error) {
	return s.repository.AssignPermissionToRole(roleId, permissionId)
}

func (s *service) ListPermissions() ([]Permission, error) {
	return s.repository.ListPermissions()
}

func (s *service) CreatePermission(p Permission) (int64, error) {
	return s.repository.CreatePermission(p)
}

func (s *service) DeletePermission(id int64) (int64, error) {
	return s.repository.DeletePermission(id)
}

func (s *service) GetPermissionsById(id int64) (Permission, error) {
	return s.repository.GetPermissionsById(id)
}

func (s *service) UpdatePermission(p Permission) (Permission, error) {
	return s.repository.UpdatePermission(p)
}
