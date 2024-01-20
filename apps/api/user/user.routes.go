package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service UserService, ctx context.Context) {
	app.Post("/user", HandleCreateUser(service, ctx))
	app.Get("/user/:id", HandleGetUser(service, ctx))
}
