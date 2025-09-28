package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		mux.ServeHTTP(w, r)
		log.Printf("%s %s | %s", r.Method, r.URL.Path, time.Since(start))
	})
}
