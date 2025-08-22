package routes

import (
	"github.com/DiansSopandi/goride_be/http/handler/v1"
	"github.com/DiansSopandi/goride_be/pkg"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// appPath := pkg.GetEnv("APP_PATH")
	appPath := pkg.Cfg.Application.AppPath
	api := app.Group(appPath)
	auth := api.Group("/auth")
	health := api.Group("/health")

	health.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"ok": true})
	})

	handler.RootHandler(api)
	handler.RolesRoutes(api)
	handler.UserRoutes(api)
	handler.AuthRoutes(auth)

	// Route untuk favicon.ico
	// app.Static("/favicon.ico", "./public/favicon.ico")
}
