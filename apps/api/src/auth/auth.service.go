package auth

import (
	"api/config"
	user "api/src/user"
	utils "api/utils"
	"context"
	"errors"
	"time"

	"log"

	"github.com/golang-jwt/jwt/v5"

	"github.com/uptrace/bun"
)

type AuthService interface {
	LogIn(user user.User) (string, error)
	Register(userParams user.User) (int64, error)
}

type service struct {
	repository user.UserRepository
	ctx        context.Context
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

func (s *service) LogIn(userParams user.User) (string, error) {
	u, err := s.repository.GetUserByName(userParams.Name, s.ctx)

	if err != nil {
		log.Fatal(err)
	}

	match, err := utils.ComparePasswordAndHash(userParams.Password, u.Password)

	if err != nil {
		log.Fatal(err)
	}

	if match {
		claims := jwt.MapClaims{
			"user_id": u.ID,
			"name":    u.Name,
			"exp":     time.Now().Add(time.Hour * 12).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		jwtToken, err := token.SignedString([]byte(config.Config("SECRET")))

		if err != nil {
			log.Fatal(err)

			return "", err
		}

		return jwtToken, nil
	}

	return "Error", ErrInvalidCredentials
}

func (s *service) Register(userParams user.User) (int64, error) {
	return s.repository.CreateUser(userParams, s.ctx)
}
