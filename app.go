package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "hxgo-skeleton/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("SERVER_HOST")
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	addr := fmt.Sprintf("http://%s%s", host, port)
	routes := server.Routes()
	server := &http.Server{
		Addr:    port,
		Handler: routes,
	}

	go func() {
		fmt.Printf("Listening on address: %s\n", addr)
		server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v", err)
	}

	fmt.Println("\nServer shutdown gracefully")
}
