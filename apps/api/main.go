package main

import (
	"api/concerns/auth"
	logging "api/concerns/logging"
	"api/database"
	"api/src/permission"
	"api/src/user"
	"context"
	"time"

	"github.com/alexlast/bunzap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	app := fiber.New()

	logging.InitLogger()

	app.Use(cors.New())

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.SendString("Online")
	})

	database := database.GetDatabase()

	database.AddQueryHook(
		bunzap.NewQueryHook(
			bunzap.QueryHookOptions{
				Logger:       zap.L(),
				SlowDuration: 200 * time.Millisecond, // Omit to log all operations as debug
			}))

	api := app.Group("/api")

	userService := user.NewService(database)
	authService := auth.NewService(database, ctx)
	permissionService := permission.NewService(database, ctx)

	user.UserRouter(api, userService, ctx)
	auth.AuthRouter(api, authService)
	permission.Router(api, permissionService)

	app.Listen(":3000")

}
