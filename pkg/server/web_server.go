package server

import (
	"net/http"
	// "html/template"
)

// func (w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello")
// 	t, _ := template.ParseFiles("index.html")
// 	t.Execute(w)
// }

var file string

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, file)
}

func Run_file(fileName string) {
	file = fileName
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
