package user

import (
	"api/middleware"
	"context"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service UserService, ctx context.Context) {
	app.Get("/user/:id", middleware.Protected(), HandleGetUserById(service, ctx))
}
