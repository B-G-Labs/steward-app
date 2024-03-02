package role

import "github.com/gofiber/fiber/v2"

func Router(app fiber.Router, service RoleService) {
	app.Post("/roles", HandleCreateRole(service))
	app.Get("/roles", HandleListRoles(service))
	app.Delete("/roles", HandleDeleteRole(service))
	app.Put("/roles", HandleUpdateRole(service))
	app.Get("/roles/:id", HandleGetRoleByID(service))
}