package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raianmr/learn-microservices-with-go/geddit/api/handlers"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.ReadRoot)

	r.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", handlers.ReadPost)
		r.Get("/", handlers.ReadPosts)
		r.Post("/", handlers.CreatePost)
		r.Put("/{id}", handlers.UpdatePost)
		r.Delete("/{id}", handlers.DeletePost)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", handlers.ReadUser)
		r.Get("/", handlers.ReadUsers)
		r.Post("/", handlers.CreateUser)
		r.Put("/{id}", handlers.UpdateUser)
		r.Delete("/{id}", handlers.DeleteUser)
	})

	r.Route("/votes", func(r chi.Router) {
		r.Get("/{id}", handlers.ReadVote)
		r.Get("/", handlers.ReadVotes)
		r.Post("/", handlers.CreateVote)
		r.Put("/{id}", handlers.UpdateVote)
		r.Delete("/{id}", handlers.DeleteVote)
	})

	return r
}
