package main

import (
	"log"
	"net/http"

	"github.com/mnlprz/buckethead-discography/database"
	"github.com/mnlprz/buckethead-discography/handlers"
)

func main() {
	db := database.GetConnectionDB()
	database.SetUpDB(db)
	r := handlers.SetUpHandlers()
	if err := http.ListenAndServe(":5555", r); err != nil {
		log.Fatal(err)
	}
}
