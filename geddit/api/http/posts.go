package http

import (
	"net/http"
)

func (s *State) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreatePost: " + s.Msg))
}

func (s *State) ReadPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) ReadPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
