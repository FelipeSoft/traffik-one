package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FelipeSoft/traffik-one/internal/core/entity"
	"github.com/FelipeSoft/traffik-one/internal/port/algorithm"
)

func StartHttpLoadBalancer(ctx context.Context, configEvent *entity.ConfigEvent) {
	mux := http.NewServeMux()

	httpLoadBalancerHost := os.Getenv("HTTP_LOAD_BALANCER_HOST")
	httpLoadBalancerPort := os.Getenv("HTTP_LOAD_BALANCER_PORT")
	httpLoadBalancerBindAddress := fmt.Sprintf("%s:%s", httpLoadBalancerHost, httpLoadBalancerPort)

	algorithmFactory := algorithm.NewAlgorithmFactory(configEvent)
	algorithmStrategy, err := algorithmFactory.Create()
	if err != nil {
		log.Fatalf("Algorithm Factory error: %v", err)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		algorithmStrategy.ReverseProxy(w, r)
	})

	server := &http.Server{
		Addr:    httpLoadBalancerBindAddress,
		Handler: mux,
	}

	go func() {
		log.Printf("[HTTP Load Balancer] Listening on %s", httpLoadBalancerBindAddress)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start the HTTP server on %s caused by error: %v", httpLoadBalancerBindAddress, err)
		}
	}()

	<-ctx.Done()
	log.Print("[HTTP Load Balancer] Shutting down...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("[HTTP Load Balancer] Shutdown failed: %v", err)
	}
}
