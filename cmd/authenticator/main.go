package main

import (
	"log"

	"github.com/FelipeSoft/traffik-one/internal/port/jsonwebtoken"
)

func main() {
	tokenManager := jsonwebtoken.NewJsonWebTokenManager()
	token, err := tokenManager.Sign(map[string]any{
		"UserId": "1",
	})
	if err != nil {
		log.Fatalf("Error during the token signing: %v", err)
	}
	log.Print(token)
}
