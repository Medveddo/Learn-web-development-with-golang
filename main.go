package main

import (
	"MyWebApp/views"
	"net/http"

	"github.com/gorilla/mux/"
)

var (
	homeView    *views.View
	contactView *views.View
)

/*
	Block of handle function. Each of them give specific page.
*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Creating our views
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
