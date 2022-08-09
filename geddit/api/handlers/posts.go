package handlers

import "net/http"

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
