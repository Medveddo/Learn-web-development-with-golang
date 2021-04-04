package main

import (
	"fmt"
	"learn-web-dev-with-go/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111"
	dbname   = "mywebapp_dev"
)

func main() {
	// Creating our database info string which will be passed to gorm.Open function and has all information needed to get successful connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	user := models.User{
		Name:  "Michael Scott",
		Email: "miScott@gmail.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	// user.Email = "mich@mail.ru"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }

	// userByEmail, err := us.ByEmail("mich@mail.ru")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userByEmail)

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)
}
