package main

import "fmt"

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

type Husky struct {
	Speaker // Embedding
}

type SpeakerPrefixer struct {
	Speaker // Chaining
}

func (sp SpeakerPrefixer) Speak() {
	fmt.Print("Prefix: ")
	sp.Speaker.Speak()
}

type Speaker interface {
	Speak()
}

func main() {
	h := Husky{SpeakerPrefixer{Cat{}}}
	h.Speak() // equals to h.Dog.Speak()
	h1 := Husky{SpeakerPrefixer{Dog{}}}
	h1.Speak()
}
