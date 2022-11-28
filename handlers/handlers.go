package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetUpHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/albums", AlbumsHandler)
	r.HandleFunc("/album/{id}", AlbumHandler)
	return r
}

func AlbumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AlbumsHandler"))
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AlbumHandler"))
}
