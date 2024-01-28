package auth

import (
	user "api/src/user"
	"context"

	"github.com/gofiber/fiber/v2"
)

func HandleLogin(service AuthService, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.User

		if err := c.BodyParser(&requestBody); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token, err := service.LogIn(requestBody, ctx)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"status": true, "message": "Success login", "data": token})
	}
}
