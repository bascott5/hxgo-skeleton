package handlers

import (
	"log"
	"net/http"

	web "hxgo-skeleton/cmd/web"
)

func IndexHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := web.EmbeddedFS.ReadFile("views/index.html")
		if err != nil {
			log.Fatal(err)
		}

		w.Write(f)
	})
}
