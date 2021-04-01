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
	// Creating our database info string which will be passed to gorm.Open function and has all information needed to get successful connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// Connecting to our database
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)        // Enable logging to see all SQL queries underlying
	db.AutoMigrate(&User{}) // Creating a users table if it doesn't exist or modify it

	// Adding some records
	// db.Create(&User{Name: "Sizikov Vitaly", Email: "sizikov.vitaly@gmail.com", Color: "blue"})
	// db.Create(&User{Name: "Dasha Malko", Email: "malkodasha25@gmail.com", Color: "green"})
	// db.Create(&User{Name: "Sergey Kolmogorov", Email: "kekivan@gmail.com", Color: "black"})

	// Querying records using Not() function
	// var users []User
	// if err := db.Not("color = ?", "green").Find(&users).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Println(users)

	// Querying records using Where() and Or() function
	// var users []User
	// if err := db.Where("name = ?", "Sizikov Vitaly").Or("color = ?", "green").Find(&users).Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Println(users)

	var users []User
	if err := db.Where("name LIKE ?", "%gey%").Find(&users).Error; err != nil {
		panic(err)
	}
	fmt.Println(users)
}
