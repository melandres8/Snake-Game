package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"log"
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
	err := http.ListenAndServe(port, route)
	if err != nil {
		log.Fatal("An error occur", err)
	}
}
