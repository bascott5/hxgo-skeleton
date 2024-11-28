package server

import (
	handlers "hxgo-skeleton/server/handlers"
	middleware "hxgo-skeleton/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", handlers.IndexHandler())

	return middleware.LoggerMiddleware(mux)
}
