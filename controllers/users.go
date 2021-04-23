package controllers

import (
	"fmt"
	"learn-web-dev-with-go/models"
	"learn-web-dev-with-go/rand"
	"learn-web-dev-with-go/views"
	"log"
	"net/http"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
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

// SignupForm is a struct that contains data
// that we can get from web form
// We use struct tags which helps us decode data using schema package
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: views.AlertMsgGeneric,
		}
		u.NewView.Render(w, vd)
		return
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		vd.Alert = &views.Alert{
			Level:   views.AlertLvlError,
			Message: err.Error(),
		}
		u.NewView.Render(w, vd)
		return
	}
	err := u.signIn(w, &user)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Used to verify the provided email address and
// password and then log the user in if they are correct
//
// POST /login
func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user, err := u.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrPasswordIncorrect:
			fmt.Fprintln(w, "Invalid password provided.")
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid email address.")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cameCookie, err := r.Cookie("came_from_url")

	// cameCookie does not exist
	if err != nil {
		//http: named cookie not present
		// Successful login
		// standart redirect
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	// cameCookie exist
	// redirect on a page where user came from
	http.Redirect(w, r, cameCookie.Value, http.StatusFound)
}

// is used to sign in the given user via cookies
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = u.us.Update(user)
		if err != nil {
			return err
		}
	}
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}

// Used to display cookies set on the current user
func (u *Users) CookieTest(w http.ResponseWriter, r *http.Request) {

	/*
		RIGHT HERE IF USER IS NOT LOGGED IN WE WANNA
		cookieCame := http.Cookie {
			Name: "came_from_url",
			Value: "/cookietest",
			HttpOnly: true,
		}
		http.SetCookie(w, &cookieCame)
		AND THEN AFTER SUCCESSFUL LOGIN WE WILL REDIRECT USER TO
		PAGE THAT CONTAINS IN THAT COOKIE
		IF ITS EMPTY WE WILL REDIRECT HIM BY DEFAULT ROUTE
	*/

	cookie, err := r.Cookie("remember_token")
	if err != nil {
		cookieCame := http.Cookie{
			Name:     "came_from_url",
			Value:    "/cookietest",
			HttpOnly: true,
		}
		http.SetCookie(w, &cookieCame)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	user, err := u.us.ByRemember(cookie.Value)
	if err != nil {
		cookieCame := http.Cookie{
			Name:     "came_from_url",
			Value:    "/cookietest",
			HttpOnly: true,
		}
		http.SetCookie(w, &cookieCame)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	fmt.Fprintln(w, user)
}
