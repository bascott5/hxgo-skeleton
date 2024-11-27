package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formCounter := r.FormValue("counter")
	counter, err := strconv.Atoi(formCounter)

	if err != nil {
		log.Fatal(err)
	}

	switch r.Pattern {
	case "/increment":
		counter++
		break
	case "/decrement":
		counter--
		break
	default:
		break
	}

	fmt.Fprint(w, counter)
}
