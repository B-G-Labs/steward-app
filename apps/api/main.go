package main

import (
	"api/concerns/auth"
	logging "api/concerns/logging"
	"api/database"
	"api/src/permission"
	"api/src/user"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type (
	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
)

func main() {
	ctx := context.Background()

	logging.InitLogger()

	database := database.GetDatabase(ctx)

	app := fiber.New(fiber.Config{
		// Global custom error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(cors.New())

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})

	api := app.Group("/api")

	userService := user.NewService(database)
	authService := auth.NewService(database, ctx)
	permissionService := permission.NewService(database, ctx)

	user.UserRouter(api, userService, ctx)
	auth.AuthRouter(api, authService)
	permission.Router(api, permissionService)

	app.Listen(":3000")

}
