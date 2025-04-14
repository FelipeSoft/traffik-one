package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/ports/http/router"
)

func StartHttpServer() {
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	httpRouter := router.NewHttpRouter()
	httpBindAddress := fmt.Sprintf("%s:%s", httpHost, httpPort)

	log.Printf("HTTP server listening on %s", httpHost)
	err := http.ListenAndServe(httpBindAddress, httpRouter)
	if err != nil {
		log.Fatalf("Could not start the HTTP server on %s caused by error: %v", httpBindAddress, err)
	}
}
