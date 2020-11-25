package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"snake-game/api/handlers"
)

func main() {
	port := ":8000"
	route := chi.NewRouter()

	route.Use(middleware.Logger)
	route.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
	})

	route.Use(cors.Handler)
	// Post a user
	route.Post("/user", handlers.PostUser())
	// Get users
	route.Get("/users", handlers.GetUsers())
	// Put to update the user score
	route.Put("/users/{id}", handlers.PutUser())

	fmt.Println("Serving on port" + port)
	_ = http.ListenAndServe(port, route)
}
