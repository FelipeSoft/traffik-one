package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/FelipeSoft/traffik-one/internal/app"
	"github.com/FelipeSoft/traffik-one/internal/bootstrap"
	"github.com/FelipeSoft/traffik-one/internal/port/bolt"
	"github.com/FelipeSoft/traffik-one/internal/port/http"
	"github.com/FelipeSoft/traffik-one/internal/port/idgen"
	"github.com/joho/godotenv"
)

func main() {
	idgen.InitNode(1)

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

	configEvent := bootstrap.LoadInitialConfig()
	appInstance := app.NewApp(ctx, configEvent)

	log.Printf("Config Event: %v", configEvent)

	go http.StartHttpServer(ctx, appInstance)
	go http.StartHttpLoadBalancer(ctx, configEvent)
	go http.StartHttpHealthChecker(ctx, configEvent, 5 * time.Second, 5)

	<-ctx.Done()
}
