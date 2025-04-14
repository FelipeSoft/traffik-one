package usecase

import "fmt"

func TestUseCase(message string) (string, error) {
	if message != "Ping" {
		return "", fmt.Errorf(`please provide the "Ping" as message param`)
	}
	return "Pong", nil
}
