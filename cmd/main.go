package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "met an error!", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Data: %s\n", data)
	})

	http.ListenAndServe(":9090", nil)
}
