package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/thyagobr/wheretogo/internal/db"
	"github.com/thyagobr/wheretogo/internal/handlers"
	"github.com/thyagobr/wheretogo/internal/middlewares"
)

func main() {
	db.InitDB()

	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3010"}, // allow all origins
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // max value not ignored by browsers
	}))

	r.Use(middlewares.HttpLoggingMiddleware)

	r.Route("/auth/login", func(r chi.Router) {
		r.Post("/", handlers.Login)
	})

	r.Route("/places", func(r chi.Router) {
		r.Get("/", handlers.GetPlaces)
		r.Get("/{id}", handlers.GetPlace)
		r.With(middlewares.AuthenticationMiddleware).Post("/", handlers.CreatePlace)
		r.Get("/{id}/events", handlers.GetPlaceEvents)
		r.Get("/search_address", handlers.SearchAddress)
		r.With(middlewares.AuthenticationMiddleware).Post("/{id}/events", handlers.CreateEvent)
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/", handlers.GetEvents)
		r.Get("/{id}", handlers.GetEvent)
		r.With(middlewares.AuthenticationMiddleware).Post("/", handlers.CreateEvent)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
