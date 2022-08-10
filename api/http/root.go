package http

import "net/http"

func (s *State) ReadRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root root"))
}
