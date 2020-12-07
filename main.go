package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
	Block of handle function. Each of them give specific page.
*/
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my cool page!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Contacts</h1>For any questions.<br><b>Phone:</b>+7-903-997-37-72<br>"+
		"<br><b>Email:</b>sizikov.vitaly@gmail.com")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	//Creating a router
	router := httprouter.New()
	// Set the handler functions to different URL's
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	// Set the custom 404 page instead of deafult.
	var h http.Handler = http.HandlerFunc(pageNotFound)
	router.NotFound = h
	log.Fatal(http.ListenAndServe(":3000", router))
}
