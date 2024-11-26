package main

import (
	"fmt"
	"net/http"
	"os"

	middleware "hxgo-skeleton/internals/middleware"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/views/index.html")
	})

	fs := http.FileServer(http.Dir("web"))
	router.Handle("/web/", http.StripPrefix("/web/", fs))

	fmt.Printf("Listening on: http://localhost:%s\n", port)
	configuredRouter := middleware.Logger(router)
	http.ListenAndServe(fmt.Sprintf(":%s", port), configuredRouter)
}
