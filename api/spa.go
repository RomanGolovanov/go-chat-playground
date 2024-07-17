package api

import (
	"net/http"
	"os"
	"path/filepath"
)

type SpaHandler struct {
	staticPath string
	indexPath  string
}

func NewSpaHandler(staticPath, indexPath string) *SpaHandler {
	return &SpaHandler{staticPath: staticPath, indexPath: indexPath}
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.staticPath, r.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
