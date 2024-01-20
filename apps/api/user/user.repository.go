package user

import (
	"api/utils"
	"context"
	"database/sql"
	"log"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

type repository interface {
	GetExistingUser(id int64, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (sql.Result, error)
}

func NewUserRepo(db *bun.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user User, ctx context.Context) (sql.Result, error) {
	hash, err := utils.GenerateFromPassword(user.Password)

	if err != nil {
		return nil, err
	}

	a, err := r.db.NewInsert().Model(&User{
		Name:     user.Name,
		Password: hash,
	}).Exec(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (r *UserRepository) GetExistingUser(id int64, ctx context.Context) (User, error) {
	var user User

	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)

	return user, err
}
