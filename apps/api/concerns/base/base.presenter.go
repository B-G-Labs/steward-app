package base

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SuccessResponse[Data any](data Data, c *fiber.Ctx) error {
	c.Status(http.StatusOK)

	return c.JSON(
		&fiber.Map{
			"data":  data,
			"error": nil,
		},
	)
}

func ErrorResponse(err error, c *fiber.Ctx, errorList ...string) error {
	c.Status(http.StatusInternalServerError)

	zap.L().Error(err.Error())

	return c.JSON(
		&fiber.Map{
			"data":   nil,
			"error":  err.Error(),
			"errors": errorList,
		},
	)
}
