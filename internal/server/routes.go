package server

import (
	handlers "hxgo-skeleton/internal/server/handlers"
	middleware "hxgo-skeleton/internal/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("cmd/web"))
	mux.Handle("/web/", http.StripPrefix("/web/", fs))

	// pages
	mux.HandleFunc("/", handlers.IndexHandler())

	// templates
	mux.HandleFunc("/add", handlers.AddHandler())

	return middleware.LoggingMiddleware(mux)
}
