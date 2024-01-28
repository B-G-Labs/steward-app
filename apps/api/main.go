package main

import (
	"api/database"
	"api/src/auth"
	"api/src/permission"
	"api/src/user"
	"context"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx := context.Background()

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(cors.New())

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})

	database := database.GetDatabase()

	api := app.Group("/api")

	userService := user.NewService(database)
	authService := auth.NewService(database)
	permissionService := permission.NewService(database, ctx)

	user.UserRouter(api, userService, ctx)
	auth.AuthRouter(api, authService, ctx)
	permission.Router(api, permissionService)

	app.Listen(":3000")

}
