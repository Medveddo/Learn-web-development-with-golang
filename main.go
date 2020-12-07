package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux/"
)

/*
	Block of handle function. Each of them give specific page.
*/
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my cool page!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contacts</h1>For any questions.<br><b>Phone:</b>+7-903-997-37-72<br>"+
		"<br><b>Email:</b>sizikov.vitaly@gmail.com")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1><ul><li><b>What a hell are doing?</b><p>"+
		"I'm try to build a web application usign Golang.</p></li><li><b>Is that legal?</b>"+
		"<p>Absolutely.</p></li></ul>")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We could not find the page you "+
		"were looking for :(</h1>"+
		"<p>Please email me if you keep being sent to an "+
		"invalid page </p>")
}

func main() {
	r := mux.NewRouter()
	// Set the custom 404 page instead of deafult.
	var h http.Handler = http.HandlerFunc(pageNotFound)
	r.NotFoundHandler = h
	// Set the handler functions to different URL's
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
