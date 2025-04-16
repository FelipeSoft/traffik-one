package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
	"github.com/FelipeSoft/traffik-one/internal/port/http"
	"github.com/joho/godotenv"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()

	err := godotenv.Load("./../.env")
	if err != nil {
		log.Fatalf("Could not load the environment variables file (.env) caused by error: %v", err)
	}

	if err := bolt.Init(os.Getenv("BOLTDB_PATH"), os.Getenv("BOLTDB_DATABASE")); err != nil {
		log.Fatalf("Failed to initialize BoltDB: %v", err)
	}
	defer bolt.Close()

	appInstance := app.NewApp()
	http.StartHttpServer(ctx, appInstance)

	<-ctx.Done()
	log.Print("Load balancer exited")
}