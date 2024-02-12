package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(app fiber.Router, service AuthService) {
	app.Post("/auth/login", HandleLogin(service))
	app.Post("/auth/register", HandleRegister(service))
}
