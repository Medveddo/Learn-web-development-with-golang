package controllers

import "MyWebApp/views"

type Static struct {
	Home    *views.View
	Contact *views.View
	FAQ     *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
		FAQ:     views.NewView("coverLayout", "views/static/FAQ.gohtml"),
	}
}
