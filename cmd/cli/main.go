package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/FelipeSoft/traffik-one/cli"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	defer stop()

	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatalf("Could not load the environment variables file (.env) caused by error: %v", err)
	}

	fmt.Println("Welcome to TraffikOne CLI!")
	fmt.Println("Type a command or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			fmt.Print("traffikone>")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("[TraffikOne Error] Error reading input:", err)
				continue
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if line == "exit" {
				cancel()
				fmt.Println("Bye!")
				break
			}

			args := strings.Fields(line)
			cli.RootCmd.SetArgs(args)
			err = cli.RootCmd.Execute()
			if err != nil {
				fmt.Printf("[FileSync Error] Command error: %s\n", err.Error())
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Exited")
	stop()
}