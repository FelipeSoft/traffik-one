package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/FelipeSoft/traffik-one/internal/ports/http"
	"github.com/joho/godotenv"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()

	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatalf("Could not load the environment variables file (.env) caused by error: %v", err)
	}

	go http.StartHttpServer()

	<-ctx.Done()
	log.Print("Load balancer exited")
}
