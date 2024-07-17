package middleware

import "net/http"

type CorsOptions struct {
	AllowOrigin      string
	AllowHeaders     string
	AllowCredentials string
	AllowMethods     string
}

func Cors(next http.Handler, o CorsOptions) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", o.AllowOrigin)
		w.Header().Set("Access-Control-Allow-Headers", o.AllowHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", o.AllowCredentials)
		w.Header().Set("Access-Control-Allow-Methods", o.AllowMethods)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
