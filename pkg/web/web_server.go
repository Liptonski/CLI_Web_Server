package web

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

func Run_file() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
