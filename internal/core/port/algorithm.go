package port

import (
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
)

type Algorithm interface {
	ReverseProxy(w http.ResponseWriter, r *http.Request)
	Next() *entity.Backend
}
