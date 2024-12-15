package server

import (
	web "hxgo-skeleton/cmd/web"
	handlers "hxgo-skeleton/internal/server/handlers"
	middleware "hxgo-skeleton/internal/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/web/", http.StripPrefix("/web/", http.FileServerFS(web.EmbeddedFS)))

	// pages
	mux.HandleFunc("/", handlers.IndexHandler())

	// templates
	mux.HandleFunc("/add", handlers.AddHandler())

	return middleware.LoggingMiddleware(mux)
}
