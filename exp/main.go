package main

import (
	"html/template"
	"os"
)

type Data struct {
	Name    string
	Boolean bool
}

func main() {
	// Get and parse our template file
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	var data Data = Data{
		Name:    "Vitaly",
		Boolean: true,
	}

	//Execute our template with some data
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	/*
		<h1>Welcome to my website!</h1>
		Glad to see you, Vitaly!
	*/

	data = Data{
		Name:    "Unknown",
		Boolean: false,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	/*
		<h1>Welcome to my website!</h1>
		If you want to save your experience, please <b><a href="#">LOG IN</a></b>
	*/
}
