package handlers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func ReadUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("it works!"))
}
