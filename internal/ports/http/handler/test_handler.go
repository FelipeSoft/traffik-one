package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/FelipeSoft/traffik-one/internal/app/usecase"
	"github.com/FelipeSoft/traffik-one/internal/domain"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid request body: %v", err)))
		return
	}

	var testDomain domain.TestDomain
	if err := json.Unmarshal(body, &testDomain); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid JSON: %v", err)))
		return
	}

	message, err := usecase.TestUseCase(testDomain.Message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Could not execute the usecase.TestUseCase: %v", err)))
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
