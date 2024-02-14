package auth

import (
	utils "api/concerns/utils"
	"api/config"
	user "api/src/user"
	"context"
	"errors"
	"fmt"
	"time"

	"log"

	"github.com/golang-jwt/jwt/v5"

	"github.com/uptrace/bun"
)

type AuthService interface {
	LogIn(user user.User, ctx context.Context) (string, error)
}

type service struct {
	repository user.UserRepository
}

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

func NewService(database *bun.DB) AuthService {
	userRepo := user.NewUserRepo(database)

	return &service{
		repository: userRepo,
	}
}

func (s *service) LogIn(userParams user.User, ctx context.Context) (string, error) {
	u, err := s.repository.GetUserByName(userParams.Name, ctx)

	fmt.Println(u)

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
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
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

	return "0", ErrInvalidCredentials
}
