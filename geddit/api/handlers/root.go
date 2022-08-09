package handlers

import "net/http"

func ReadRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
