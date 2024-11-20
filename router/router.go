package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/cloak/domain/auth"
)

func Run(app *fiber.App) {
	// Health check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	auth.Handler.Route(
		app.Group("/auth"),
	)

	// Group for API routes
	api := app.Group("/api")
	{
		api.Use(auth.Middleware)

		// Version 1 of the API
		v1 := api.Group("/v1")
		{
			v1.Get("/docs", func(c *fiber.Ctx) error {
				return c.SendString("hello")
			})
		}
	}
}
