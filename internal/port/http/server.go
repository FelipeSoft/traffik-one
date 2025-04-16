package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/port/http/router"
)

func StartHttpServer(ctx context.Context, app *app.App) {
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	httpRouter := router.NewHttpRouter(app)
	httpBindAddress := fmt.Sprintf("%s:%s", httpHost, httpPort)

	server := &http.Server{
		Addr:    httpBindAddress,
		Handler: httpRouter,
	}

	go func() {
		log.Printf("HTTP server listening on %s", httpHost)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start the HTTP server on %s caused by error: %v", httpBindAddress, err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down HTTP server...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
}

