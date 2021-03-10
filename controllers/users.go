package controllers

import (
	"MyWebApp/views"
	"fmt"
	"net/http"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can
// create a new user account
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// r.PostForm = map[string] []strings
	fmt.Fprintln(w, r.PostForm["email"]) //slice
	//fmt.Fprintln(w, r.PostFormValue("email"))    //value
	fmt.Fprintln(w, r.PostForm["password"]) //slice
	//fmt.Fprintln(w, r.PostFormValue("password")) //value
}
