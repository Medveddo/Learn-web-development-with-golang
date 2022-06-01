package main

import (
	"fmt"
	"learnwebdev/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111"
	dbname   = "mywebapp_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()
	// us.AutoMigrate()

	user := models.User{
		Name:     "Sizikov Vitaly",
		Email:    "sizikov.vitaly1@gmail.com",
		Password: "111",
		Remember: "abc123",
	}
	err = us.Create(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", user)
	user2, err := us.ByRemember("abc123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)
}
