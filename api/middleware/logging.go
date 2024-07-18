package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func NewLoggingMiddleware() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			wrappedWriter := &wrappedResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(wrappedWriter, r)
			log.Println(wrappedWriter.statusCode, r.Method, r.URL.Path, time.Since(start))
		})
	}
}
