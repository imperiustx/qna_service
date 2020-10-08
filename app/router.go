package app

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

		r.Post("/admins", a.handleAdminCreate)
		r.Put("/admins/{id}", a.handleAdminUpdate)

		r.Post("/users", a.handleUserCreate)
		r.Put("/users/{id}", a.handleUserUpdate)
		r.Get("/users/{id}", a.handleUserFind)

		r.Post("/questions", a.handleQuestionCreate)
		r.Get("/questions", a.handleGetAllQuestions)
		r.Get("/questions/open", a.handleGetAllOpenQuestions)
		r.Put("/questions/{id}", a.handleQuestionUpdate)
		r.Get("/questions/{id}", a.handleViewQuestionAndAnswer)
		r.Put("/questions/{id}/up", a.handleQuestionVoteUp)
		r.Put("/questions/{id}/down", a.handleQuestionVoteDown)

		r.Post("/answers", a.handleAnswerCreate)
		r.Put("/answers/{id}", a.handleAnswerUpdate)
		r.Put("/answers/{id}/up", a.handleAnswerVoteUp)
		r.Put("/answers/{id}/down", a.handleAnswerVoteDown)

	})
	a.mux.NotFound(a.handleResourceNotFound)
}
