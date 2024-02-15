package auth

import (
	utils "api/concerns/utils"
	"api/concerns/validation"
	"api/config"
	user "api/src/user"

	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/uptrace/bun"
)

type AuthService interface {
	LogIn(user user.User) (JWTResult, error)
	Register(userParams user.User) (int64, error)
}

type service struct {
	repository user.UserRepository
	ctx        context.Context
}

type JWTResult struct {
	token          string
	expiryDateUnix int64
}

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

func NewService(database *bun.DB, ctx context.Context) AuthService {
	userRepo := user.NewUserRepo(database)

	return &service{
		repository: userRepo,
		ctx:        ctx,
	}
}

func (s *service) LogIn(userParams user.User) (JWTResult, error) {
	result := JWTResult{}

	validator := validation.Validate()

	err := validator.Struct(userParams)

	if err != nil {
		return result, err
	}

	user, err := s.repository.GetUserByName(userParams.Name, s.ctx)

	if err != nil {
		return result, err
	}

	match, err := utils.ComparePasswordAndHash(userParams.Password, user.Password)

	if err != nil {
		return result, err
	}

	if match {
		expiry := time.Now().Add(time.Hour * 12).Unix()

		claims := jwt.MapClaims{
			"user_id": user.ID,
			"name":    user.Name,
			"exp":     expiry,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		jwtToken, err := token.SignedString([]byte(config.Config("SECRET")))

		result.token = jwtToken
		result.expiryDateUnix = expiry

		if err != nil {
			return result, err
		}

		return result, nil
	}

	return result, ErrInvalidCredentials
}

func (s *service) Register(userParams user.User) (int64, error) {
	validator := validation.Validate()

	if err := validator.Struct(userParams); err != nil {
		return 0, err
	}

	result, err := s.repository.CreateUser(userParams, s.ctx)

	return result, err
}
