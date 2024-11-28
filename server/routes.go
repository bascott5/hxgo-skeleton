package server

import (
	handlers "hxgo-skeleton/server/handlers"
	middleware "hxgo-skeleton/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler())

	return middleware.LoggerMiddleware(mux)
}
