package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(t)
		fmt.Printf(
			"%d-%d-%d %d:%d:%d "+`"%s %s"`+" in %s\n",
			t.Year(),
			t.Month(),
			t.Day(),
			t.Hour(),
			t.Minute(),
			t.Second(),
			r.Method,
			r.Pattern,
			duration)
	})
}
