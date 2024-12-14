package server

import (
	handlers "hxgo-skeleton/server/handlers"
	middleware "hxgo-skeleton/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("client"))
	mux.Handle("/client/", http.StripPrefix("/client/", fs))

	// pages
	mux.HandleFunc("/", handlers.IndexHandler())

	// templates
	mux.HandleFunc("/add", handlers.AddHandler())

	return middleware.LoggingMiddleware(mux)
}
