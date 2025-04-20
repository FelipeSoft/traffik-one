package middleware

import (
	"net/http"
	"strings"

	"github.com/FelipeSoft/traffik-one/internal/core/port"
)

type AuthenticationMiddleware struct {
	tokenManager port.TokenManager
}

func NewAuthenticationMiddleware(tokenManager port.TokenManager) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthenticationMiddleware) HandleBearerToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: missing authorization header"))
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: invalid authorization header format"))
			return
		}

		token := parts[1]
		ok, err := m.tokenManager.Verify(token)
		
		if !ok || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: invalid bearer token"))
			return
		}

		next(w, r)
	}
}
