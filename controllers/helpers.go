package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// ParseForm function can be used multiple times to parse input data
// from http.Request to a special struct using schema package
func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}
