package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/thyagobr/wheretogo/internal/handlers"
  "github.com/thyagobr/wheretogo/internal/db"

)

func main() {
	db.InitDB()

	r := chi.NewRouter()

	r.Route("/places", func(r chi.Router) {
		r.Get("/", handlers.GetPlaces)
		r.Get("/{id}", handlers.GetPlace)
		r.Post("/", handlers.CreatePlace)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
