package main

import (
	"fmt"
	"net/http"
	"os"

	server "hxgo-skeleton/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("SERVER_HOST")
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	addr := fmt.Sprintf("http://%s%s", host, port)

	router := server.Routes()

	fmt.Printf("Listening on address: %s\n", addr)
	http.ListenAndServe(port, router)
}
