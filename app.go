package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
		fmt.Printf("%s %q\n", r.Method, r.Pattern)
	})
}

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
	configuredRouter := logger(router)
	http.ListenAndServe(fmt.Sprintf(":%s", port), configuredRouter)
}
