package permission

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleListPermission(service PermissionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := service.ListPermissions()

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   data,
			"error":  err,
			"body":   nil,
		})
	}
}

func HandleCreatePermission(service PermissionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Permission

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

		dataID, err := service.CreatePermission(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   dataID,
			"error":  err,
		})
	}
}

func HandleUpdatePermission(service PermissionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Permission

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
		paramId, err := c.ParamsInt("id")

		requestBody.ID = int64(paramId)
		requestBody.UpdatedAt = time.Now()

		dataID, err := service.UpdatePermission(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   dataID,
			"error":  err,
		})
	}
}

func HandleDeletePermission(service PermissionService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		paramId, err := c.ParamsInt("id")

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		dataID, err := service.DeletePermission(int64(paramId))

		if err != nil {
			c.Status(http.StatusBadRequest)

			response := &fiber.Map{
				"status": false,
				"data":   "",
				"error":  err.Error(),
			}

			return c.JSON(response)
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   dataID,
			"error":  err,
		})
	}
}
