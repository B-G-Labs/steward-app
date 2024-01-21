package user

import (
	"context"
	"errors"

	"github.com/uptrace/bun"
)

type UserService interface {
	GetUserByName(name string, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (int64, error)
	GetUserById(id int64, ctx context.Context) (User, error)
}

type service struct {
	repository UserRepository
}

var (
	ErrUserDoesntExist = errors.New("No user found")
)

func NewService(database *bun.DB) UserService {
	userRepo := NewUserRepo(database)

	return &service{
		repository: userRepo,
	}
}

func (s *service) CreateUser(user User, ctx context.Context) (int64, error) {
	return s.repository.CreateUser(user, ctx)
}

func (s *service) GetUserByName(name string, ctx context.Context) (User, error) {
	return s.repository.GetUserByName(name, ctx)
}
func (s *service) GetUserById(id int64, ctx context.Context) (User, error) {
	return s.repository.GetUserById(id, ctx)
}
