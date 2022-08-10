package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func DefaultRouter(s *State) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", s.ReadRoot)

	r.Route("/posts", func(r chi.Router) {
		r.Get("/{id}", s.ReadPost)
		r.Get("/", s.ReadPosts)
		r.Post("/", s.CreatePost)
		r.Put("/{id}", s.UpdatePost)
		r.Delete("/{id}", s.DeletePost)
	})

	// r.Route("/users", func(r chi.Router) {
	// 	r.Get("/{id}", ReadUser)
	// 	r.Get("/", ReadUsers)
	// 	r.Post("/", CreateUser)
	// 	r.Put("/{id}", UpdateUser)
	// 	r.Delete("/{id}", DeleteUser)
	// })

	// r.Route("/votes", func(r chi.Router) {
	// 	r.Get("/{id}", ReadVote)
	// 	r.Get("/", ReadVotes)
	// 	r.Post("/", CreateVote)
	// 	r.Put("/{id}", UpdateVote)
	// 	r.Delete("/{id}", DeleteVote)
	// })

	return r
}
