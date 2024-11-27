package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "hxgo-skeleton/internals/handlers"
	middleware "hxgo-skeleton/internals/middleware"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	router := http.NewServeMux()

	// pages
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/views/index.html")
	})

	// templates
	router.HandleFunc("/increment", handlers.CounterHandler)
	router.HandleFunc("/decrement", handlers.CounterHandler)

	fs := http.FileServer(http.Dir("web"))
	router.Handle("/web/", http.StripPrefix("/web/", fs))

	configuredRouter := middleware.LoggerMiddleware(router)
	fmt.Printf("Listening on: http://localhost:%s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), configuredRouter)
}
