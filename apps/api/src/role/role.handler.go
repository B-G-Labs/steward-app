package role

import (
	presenter "api/concerns/base"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleGetRoleByID(service RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Role

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		role, err := service.GetRoleByID(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(role, c)
	}
}

func HandleCreateRole(service RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Role

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		dataID, err := service.CreateRole(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(dataID, c)
	}
}

func HandleListRoles(service RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, err := service.ListRoles()

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(roles, c)
	}
}

func HandleDeleteRole(service RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Role

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		deleted, err := service.DeleteRole(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(deleted, c)
	}
}

func HandleUpdateRole(service RoleService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody Role

		err := c.BodyParser(&requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		updated, err := service.UpdateRole(requestBody)

		if err != nil {
			c.Status(http.StatusBadRequest)

			return presenter.ErrorResponse(err, c)
		}

		return presenter.SuccessResponse(updated, c)
	}
}


