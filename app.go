package cloak

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/cloak/config"
	"github.com/avila-r/cloak/router"
)

var (
	app = func() *fiber.App {
		new := fiber.New()

		router.Run(new)

		return new
	}()
)

func Start() {
	if err := app.Listen(config.Properties.ServerURL); err != nil {
		log.Fatalf("failed to start http server - %v", err.Error())
	}
}
