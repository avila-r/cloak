package auth

import "github.com/Nerzal/gocloak"

type AuthProvider struct {
	Gocloak      gocloak.GoCloak
	ClientID     string
	ClientSecret string
	Realm        string
}

type LoginRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
