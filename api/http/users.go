package http

import (
	"net/http"
)

func (s *State) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("CreateUser: " + s.Msg))
}

func (s *State) ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func (s *State) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
