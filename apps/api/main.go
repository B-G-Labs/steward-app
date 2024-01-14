package main

import (
	// "api/database"
	"api/database"
	"api/user"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx := context.Background()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	database := database.GetDatabase()

	api := app.Group("/api")

	userService := user.NewService(database)

	user.UserRouter(api, userService, ctx)

	app.Listen(":3000")

}
