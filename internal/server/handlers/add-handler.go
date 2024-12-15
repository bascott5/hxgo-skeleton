package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func AddHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}

		note := r.FormValue("note")

		tmpl, err := template.New("item.tmpl").ParseFiles("web/views/hx-templates/item.tmpl")
		if err != nil {
			log.Fatalln(err)
		}

		err = tmpl.Execute(w, note)
		if err != nil {
			log.Fatalln(err)
		}
	})
}
