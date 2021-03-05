package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux/"
)

var (
	homeTemplate    *template.Template
	contactTemplate *template.Template
)

/*
	Block of handle function. Each of them give specific page.
*/

func home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
