package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111"
	dbname   = "mywebapp_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
	Color string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{})
	var u User

	// EXAMPLE 1
	// db = db.Where("email = ?", "blah@blah.com").First(&u)
	// if db.Error != nil {
	// 	panic(db.Error)
	// }

	// EXAMPLE 2
	// if err := db.Where("email = ?", "blah@blah.com").First(&u).Error; err != nil {
	// 	panic(err)
	// }

	// EXAMPLE 3
	// db = db.Where("email = ?", "blah@blah.com").First(&u)
	// errors := db.GetErrors()
	// if len(errors) > 0 {
	// 	fmt.Println(errors)
	// 	os.Exit(1)
	// }

	// EXAMPLE 4
	// db = db.Where("email = ?", "blah@blah.com").First(&u)
	// if db.RecordNotFound() {
	// 	fmt.Println("No user found")
	// } else if db.Error != nil {
	// 	panic(db.Error)
	// } else {
	// 	fmt.Println(u)
	// }

	// EXAMPLE 5
	if err := db.Where("email = ?", "blah@blah.com").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found")
		default:
			panic(err)
		}
	}
	fmt.Println(u)

}
