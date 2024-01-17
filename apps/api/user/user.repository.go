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
	GetExistingUser(id int, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (sql.Result, error)
}

func NewUserRepo(db *bun.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user User, ctx context.Context) (sql.Result, error) {

	p := &utils.ArgonParams{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	hash, err := utils.GenerateFromPassword(user.Password, p)

	if err != nil {
		return nil, err
	}

	a, err := r.db.NewInsert().Model(&User{
		ID:       user.ID,
		Name:     user.Name,
		Password: hash,
	}).Exec(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (r *UserRepository) GetExistingUser(id int, ctx context.Context) (User, error) {
	var user User

	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)

	return user, err
}
