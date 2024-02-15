package auth

import (
	presenter "api/concerns/base"
	utils "api/concerns/utils"
	user "api/src/user"

	"github.com/gofiber/fiber/v2"
)

func HandleLogin(service AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody user.User

		if err := c.BodyParser(&requestBody); err != nil {
			return presenter.ErrorResponse(err, c)
		}

		jwtResult, err := service.LogIn(requestBody)

		if err != nil {
			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(&fiber.Map{
			"token":       jwtResult.token,
			"expiry_data": jwtResult.expiryDateUnix,
		}, c)
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
			return presenter.ErrorResponse(err, c, utils.ValidationErrorToStrList(err)...)
		}

		return c.JSON(fiber.Map{"data": id})
	}
}
