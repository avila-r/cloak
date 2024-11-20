package auth

import (
	"github.com/gofiber/fiber/v2"
)

var (
	Middleware = func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()["Authorization"]

		var token string
		if len(header) > 0 {
			token = header[0]
		}

		result, err := Handler.Keycloak.Gocloak.RetrospectToken(
			token,
			Handler.Keycloak.ClientID,
			Handler.Keycloak.ClientSecret,
			Handler.Keycloak.Realm)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or malformed token",
			})
		}

		if !result.Active {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		return nil
	}
)
