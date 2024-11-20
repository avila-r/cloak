package auth

import (
	"github.com/Nerzal/gocloak"
	"github.com/gofiber/fiber/v2"

	"github.com/avila-r/cloak/config"
)

type handler struct {
	Keycloak AuthProvider
}

var (
	Handler = &handler{
		Keycloak: AuthProvider{
			Gocloak:      gocloak.NewClient(config.Properties.KeycloakRealm),
			ClientID:     config.Properties.KeycloakClient,
			ClientSecret: config.Properties.KeycloakClientSecret,
			Realm:        config.Properties.KeycloakRealm,
		},
	}
)

func (h *handler) Route(r fiber.Router) {
	r.Post("/login", func(c *fiber.Ctx) error {
		request := LoginRequest{}
		if err := c.BodyParser(&request); err != nil {
			return err
		}

		jwt, err := h.Keycloak.Gocloak.Login(
			h.Keycloak.ClientID,
			h.Keycloak.ClientSecret,
			h.Keycloak.Realm,
			request.Username,
			request.Password,
		)

		if err != nil {
			return err
		}

		response := LoginResponse{
			AccessToken:  jwt.AccessToken,
			RefreshToken: jwt.RefreshToken,
			ExpiresIn:    jwt.ExpiresIn,
		}

		return c.JSON(response)
	})
}
