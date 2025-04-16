package middleware

import "github.com/FelipeSoft/traffik-one/internal/port/jsonwebtoken"

type Container struct {
	AuthenticationMiddleware *AuthenticationMiddleware
}

func NewContainer() *Container {
	tokenManager := jsonwebtoken.NewJsonWebTokenManager()

	return &Container{
		AuthenticationMiddleware: NewAuthenticationMiddleware(tokenManager),
	}
}
