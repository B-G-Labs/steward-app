package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(service UserService, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody User

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		result, err := service.CreateUser(requestBody, ctx)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   result,
			"error":  err,
			"body":   requestBody,
		})
	}
}

func HandleGetUser(service UserService, ctx context.Context) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := strconv.ParseInt(c.Params("id"), 10, 64)

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		result, err := service.GetExistingUser(id, ctx)

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   result,
			"error":  err,
			"body":   id,
		})
	}
}
