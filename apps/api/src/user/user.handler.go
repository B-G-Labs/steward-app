package user

import (
	presenter "api/concerns/base"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(service UserService, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody User

		err := c.BodyParser(&requestBody)

		if err != nil {
			response := presenter.ErrorResponse(err, c)

			return c.JSON(response)
		}

		result, err := service.CreateUser(requestBody, ctx)

		if err != nil {
			response := presenter.ErrorResponse(err, c)

			return c.JSON(response)
		}

		response := presenter.SuccessResponse(result, c)

		return c.JSON(response)
	}
}

func HandleGetUserById(service UserService, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)

		if err != nil {
			response := presenter.ErrorResponse(err, c)

			return c.JSON(response)
		}

		result, err := service.GetUserById(id, ctx)

		if err != nil {
			response := presenter.ErrorResponse(err, c)

			return c.JSON(response)
		}

		response := presenter.SuccessResponse(result, c)

		return response
	}
}
