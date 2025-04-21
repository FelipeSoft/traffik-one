package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()

	mux := http.NewServeMux()
	httpBindAddresses := []string{"127.0.0.1:8001"}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello from server 2")
		w.WriteHeader(http.StatusOK)
	})

	for idx, bindAddress := range httpBindAddresses {
		go func(idx int) {
			log.Printf("[HTTP Server] Listening on %s", bindAddress)
			if err := http.ListenAndServe(bindAddress, mux); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Could not start the HTTP server on %s caused by error: %v", bindAddress, err)
			}
		}(idx)
	}

	<-ctx.Done()
}
