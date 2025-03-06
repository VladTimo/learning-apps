package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	// Start a goroutine to increment the counter every 10 seconds
	go func() {
		for {
			time.Sleep(10 * time.Second)
			mu.Lock()
			counter++
			mu.Unlock()
		}
	}()

	// HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count := counter
		mu.Unlock()
		fmt.Fprintf(w, "Hello, World! Counter: %d", count)
	})

	// Start server
	port := "8089"
	fmt.Println("Server started on port", port)
	http.ListenAndServe(":"+port, nil)
}
