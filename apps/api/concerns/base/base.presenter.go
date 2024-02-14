package base

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse[Data any](data Data, c *fiber.Ctx) *fiber.Map {
	c.Status(http.StatusOK)

	return &fiber.Map{
		"data":  data,
		"error": nil,
	}
}

func ErrorResponse(err error, c *fiber.Ctx) *fiber.Map {
	c.Status(http.StatusInternalServerError)

	return &fiber.Map{
		"data":  nil,
		"error": err,
	}
}
