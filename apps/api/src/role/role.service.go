package role

import (
	"api/concerns/validation"
	"context"

	"github.com/uptrace/bun"
)

type RoleService interface {
	ListRoles() ([]Role, error)
	DeleteRole(role Role) (int64, error)
	CreateRole(role Role) (int64, error)
	UpdateRole(role Role) (Role, error)
	GetRoleByID(role Role) (Role, error)
}

type service struct {
	repository RoleRepository
	ctx        context.Context
}

func NewService(database *bun.DB, ctx context.Context) RoleService {
	roleRepo := NewRoleRepository(database, ctx)
	
	return &service{
		repository: roleRepo,
		ctx:        ctx,
	}
}

func (s *service) GetRoleByID(role Role) (Role, error) {
	if _, err := validation.ValidateEntity(role); err != nil {
		return role, err
	}
	
	return s.repository.getRoleById(role.ID)
}

func (s *service) CreateRole(role Role) (int64, error) {
	if _, err := validation.ValidateEntity(role); err != nil {
		return 0, err
	}
	
	return s.repository.CreateRole(role)
}

func (s *service) DeleteRole(role Role) (int64, error) {
	if _, err := validation.ValidateEntity(role); err != nil {
		return 0, err
	}
	
	return s.repository.DeleteRole(role)
}

func (s *service) ListRoles() ([]Role, error) {
	return s.repository.ListRoles()
}

func (s *service) UpdateRole(role Role) (Role, error) {
	if _, err := validation.ValidateEntity(role); err != nil {
		return role, err
	}
	
	return s.repository.UpdateRole(role)
}
