package main

import (
	"html/template"
	"os"
)

// Character is character struct
type Character struct {
	Name   string
	HP     uint
	Attack uint
}

func main() {
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	var data Character
	data.Name = "<script>alert('Howdy!');</script>"
	data.HP = 100
	data.Attack = 15

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
