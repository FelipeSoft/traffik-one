package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var (
		mu   sync.Mutex
		root []string
		wg   sync.WaitGroup
	)

	ch := make(chan []string)

	// Simula envio de listas atualizadas
	go func() {
		ch <- []string{"hello", "world", "!"}
		time.Sleep(3 * time.Second)
		ch <- []string{"how", "are", "you?"}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range ch {
			mu.Lock()
			root = msg
			log.Printf("Current Value: %v", root)
			mu.Unlock()
		}
	}()

	// Espera a goroutine finalizar
	wg.Wait()

	// Agora sim é seguro acessar root com lock
	mu.Lock()
	log.Printf("Updated List (final): %v", root)
	mu.Unlock()
}
