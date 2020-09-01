package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Main Page</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact Page</h1>")
	} else {
		fmt.Fprint(w, "<h1>Fuck Off!</h1>")
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlerFunc)
	r.HandleFunc("/contact", handlerFunc)

	http.ListenAndServe(":3000", r)
}
