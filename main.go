package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"Go-Web/views"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

//func notFound(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	w.WriteHeader(http.StatusNotFound)
//	fmt.Fprint(w, "<h1>Sorry, we couldn't find the page you were looking for</h1>")
//}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")

	r := mux.NewRouter()
	//r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
}
