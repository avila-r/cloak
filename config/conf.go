package config

import (
	"log"

	"github.com/joeshaw/envdecode"

	"github.com/avila-r/env"
)

type application_conf struct {
	ServerURL            string `env:"SERVER_URL,required"`
	KeycloakURL          string `env:"KEYCLOAK_URL,required"`
	KeycloakRealm        string `env:"KEYCLOAK_REALM,required"`
	KeycloakClient       string `env:"KEYCLOAK_CLIENT_ID,required"`
	KeycloakClientSecret string `env:"KEYCLOAK_CLIENT_SECRET,required"`
}

var (
	Properties = func() *application_conf {
		if !env.IsLoaded {
			env.Load()
		}

		var (
			a application_conf
		)

		if err := envdecode.StrictDecode(&a); err != nil {
			log.Fatalf("failed to decode app envs - %v", err.Error())
		}

		return &a
	}()
)
