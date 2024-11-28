package handlers

import "net/http"

func IndexHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/views/index.html")
	})
}
