package server

import (
	"fmt"
	handlers "hxgo-skeleton/server/handlers"
	middleware "hxgo-skeleton/server/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("client/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// pages
	mux.HandleFunc("/", handlers.IndexHandler())
	mux.HandleFunc("/robots.txt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "User-agent: *\nAllow: /")
	}))

	// templates
	mux.HandleFunc("/add", handlers.AddHandler())

	return middleware.LoggerMiddleware(mux)
}
