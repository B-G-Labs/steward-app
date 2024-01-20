package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service AuthService, ctx context.Context) {
	app.Post("/auth/login", HandleLogin(service, ctx))
}
