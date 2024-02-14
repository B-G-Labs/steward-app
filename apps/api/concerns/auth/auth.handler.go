package auth

import (
	user "api/src/user"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func HandleLogin(service AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.User

		if err := c.BodyParser(&requestBody); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		token, err := service.LogIn(requestBody)

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"status": true, "message": "Success login", "data": token})
	}
}

func HandleRegister(service AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.User

		if err := c.BodyParser(&requestBody); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		id, err := service.Register(requestBody)

		if err != nil {

			zap.L().Sugar().Errorf("Failed to create user %v", err)

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"status": true, "message": "Success register", "data": id})
	}
}
