package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintf(w, "Hello, World!")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	done := make(chan error)
	go func() {
		fmt.Println("server start")
		done <- server.ListenAndServe()
	}()

	select {
	case err := <-done:
		log.Fatal(err)
		os.Exit(1)
	case <-ctx.Done():
		fmt.Println("server stop")
		sCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.Shutdown(sCtx)
	}
}
