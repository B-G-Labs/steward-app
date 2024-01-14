package user

import (
	"context"
	"database/sql"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

type repository interface {
	GetExistingUser(id int, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (sql.Result, error)
}

func NewUserRepo(db *bun.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user User, ctx context.Context) (sql.Result, error) {
	a, err := r.db.NewInsert().Model(&User{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}).Exec(ctx)

	if err != nil {
		return a, err
	}

	return a, nil
}

func (r *UserRepository) GetExistingUser(id int, ctx context.Context) (User, error) {
	var user User

	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)

	return user, err
}
