package main

import (
	"fmt"
	"learn-web-dev-with-go/rand"
)

func main() {
	fmt.Println(rand.String(10))
	fmt.Println(rand.RememberToken())
}
