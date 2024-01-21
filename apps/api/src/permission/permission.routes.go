package permission

import (
	"api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app fiber.Router, service PermissionService) {
	app.Post("/permission", middleware.Protected(), HandleCreatePermission(service))
	app.Get("/permission", middleware.Protected(), HandleListPermission(service))
	app.Put("/permission/:id", middleware.Protected(), HandleUpdatePermission(service))
	app.Delete("/permission/:id", middleware.Protected(), HandleDeletePermission(service))
}
