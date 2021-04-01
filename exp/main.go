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

	// Query Raw SQL with Scan
	// type Result struct {
	// 	ID    int
	// 	Name  string
	// 	Email string
	// 	Color string
	// }
	// var result Result
	// db.Raw("SELECT id, name, email, color FROM users WHERE id = ?", 2).Scan(&result)
	// fmt.Printf("%+v\n", result)
	// {ID:2 Name:Dasha Malko Email:malkodasha25@gmail.com Color:green}

	// Query getting result as *sql.Row
	// var name, email string
	// row := db.Raw("SELECT name, email from USERS where name = ?", "Sizikov Vitaly").Row()
	// row.Scan(&name, &email)
	// fmt.Println(name, email)
	// Sizikov Vitaly sizikov.vitaly@gmail.com

	// Update Raw SQL
	db.Exec("UPDATE users SET name = ? WHERE id = ?", "Vitaly Sizikov", 1)
	var u User
	db.Where("id = ?", 1).First(&u)
	fmt.Println(u)
}
