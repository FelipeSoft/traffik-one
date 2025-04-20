package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/app"
)

func StartHttpServer(ctx context.Context, app *app.App) {
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")
	httpRouter := RegisterRoutes(app)
	httpBindAddress := fmt.Sprintf("%s:%s", httpHost, httpPort)

	server := &http.Server{
		Addr:    httpBindAddress,
		Handler: httpRouter,
	}

	go func() {
		log.Printf("[HTTP Server] Listening on %s", httpBindAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start the HTTP server on %s caused by error: %v", httpBindAddress, err)
		}
	}()

	<-ctx.Done()
	log.Print("[HTTP Server] Shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("[HTTP Server] Shutdown failed: %v", err)
	}
}
