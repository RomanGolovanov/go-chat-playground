package handlers

import "net/http"

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
