package middleware

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/port/jsonwebtoken"
)

type Container struct {
	AuthenticationMiddleware *AuthenticationMiddleware
}

func NewContainer() *Container {
	tokenManager := jsonwebtoken.NewJsonWebTokenManager()

	log.Println("[Middleware Container] Dependencies loaded successfully")

	return &Container{
		AuthenticationMiddleware: NewAuthenticationMiddleware(tokenManager),
	}
}
