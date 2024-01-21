package user

import (
	"api/utils"
	"context"

	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

type repository interface {
	GetExistingUser(id int64, ctx context.Context) (User, error)
	CreateUser(user User, ctx context.Context) (int64, error)
}

func NewUserRepo(db *bun.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user User, ctx context.Context) (int64, error) {
	hash, err := utils.GenerateFromPassword(user.Password)

	if err != nil {
		return 0, err
	}

	model := &User{
		Name:     user.Name,
		Password: hash,
		Active:   true,
	}

	sqlResult, err := r.db.NewInsert().Model(model).Exec(ctx)

	if err != nil {
		return 0, nil
	}

	return sqlResult.LastInsertId()
}

func (r *UserRepository) GetUserByName(name string, ctx context.Context) (User, error) {
	user := new(User)

	err := r.db.NewSelect().Model(user).Where("name = ?", name).Scan(ctx)

	if err != nil {
		return *user, err
	}

	return *user, nil
}

func (r *UserRepository) GetUserById(id int64, ctx context.Context) (User, error) {
	user := new(User)

	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
