package main

import (
	"log"
	"net/http"

	"Practice6/structure/db"
	"Practice6/structure/movies"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer database.Close()

	repo := movies.NewRepository(database)
	handler := movies.NewHandler(repo)

	http.HandleFunc("/movies", handler.GetMovies)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
