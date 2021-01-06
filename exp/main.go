package main

import (
	"fmt"
	"html/template"
	"math"
	"os"
)

type Inventory struct {
	Size int
}

type Character struct {
	Name    string
	HP      uint
	Attack  uint
	Boolean bool
	FLvalue float64
	Arr     []int
	Map     map[string]string
	Invent  Inventory
}

func main() {
	var pi = math.Pi
	t, err := template.ParseFiles("template.gohtml")
	if err != nil {
		panic(err)
	}

	inv := Inventory{Size: 33}
	var data Character
	data = Character{
		Name:    "Vitaly",
		HP:      100,
		Attack:  25,
		Boolean: true,
		FLvalue: pi,
		Arr:     []int{10, 20, 30},
		Map: map[string]string{
			"key1": "value1", "key2": "value2",
		},
		Invent: inv,
	}

	fmt.Println(data.Arr)

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
