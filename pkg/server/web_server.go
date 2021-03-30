package server

import (
	"net/http"
)

var file string

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, file)
}

func Run_file(fileName string) {
	file = fileName
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
