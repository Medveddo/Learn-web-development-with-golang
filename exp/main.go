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

	//userInAgeRange, err := us.InAgeRange(18, 30)
	// &[{{2 2021-04-04 18:01:19.502065 +0700 +07 2021-04-04 18:01:19.502065 +0700 +07 <nil>} Vitaly Sizikov 18 sizikov.vitaly@gmail.com}
	// {{3 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil>}  25 bob@bob.com}]

	//userInAgeRange, err := us.InAgeRange(1, 2)
	//&[]
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userInAgeRange)

}
