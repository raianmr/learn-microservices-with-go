package handlers

import "net/http"

func CreateVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadVotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func UpdateVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func DeleteVote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
