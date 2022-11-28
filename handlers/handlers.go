package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mnlprz/buckethead-discography/services"
)

func SetUpHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/albums", AlbumsHandler)
	r.HandleFunc("/album/{id}", AlbumHandler)
	r.HandleFunc("/album", NewAlbumHandler).Methods("POST")
	return r
}

func AlbumsHandler(w http.ResponseWriter, r *http.Request) {
	albums := services.GetAllAlbums()
	data, err := json.Marshal(albums)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
}

func AlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AlbumHandler"))
}

func NewAlbumHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("NewAlbum"))
}
