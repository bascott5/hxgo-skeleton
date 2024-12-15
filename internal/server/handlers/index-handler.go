package handlers

import "net/http"

func IndexHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "cmd/web/views/index.html")
	})
}
