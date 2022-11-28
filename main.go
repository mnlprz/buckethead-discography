package main

import (
	"log"
	"net/http"

	"github.com/mnlprz/buckethead-discography/handlers"
)

func main() {

	r := handlers.SetUpHandlers()
	if err := http.ListenAndServe(":5555", r); err != nil {
		log.Fatal(err)
	}
}
