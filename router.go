package main

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (a *Application) setUpMux() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           60, // Maximum value not ignored by any of major browsers
	}).Handler)
	return mux
}

func (a *Application) setUpRoute() {
	a.mux.Get("/", a.handleWelcome)
	a.mux.Get("/health", a.handleHealthCheck)
	a.mux.Route("/api/v1", func(r chi.Router) {

		// r.Post("/admins", a.auth(a.handleAdminCreate, false, false))
		// r.Post("/admins/login", a.auth(a.handleAdminLogin, false, false))
		// r.Put("/admins/{id}", a.auth(a.handleAdminUpdate, true, false))

		// r.Post("/users", a.auth(a.handleUserCreate, false, false))
		// r.Post("/users/login", a.auth(a.handleUserLogin, false, false))
		// r.Put("/users/{id}", a.auth(a.handleUserUpdate, true, false))

	})
	a.mux.NotFound(a.handleResourceNotFound)
}
