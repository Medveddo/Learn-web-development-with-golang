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
	var users []User // = User{
	// 	Color: "Red",
	// 	Email: "medveddo@gmail.com",
	// }
	//db.First(&u)
	//db.Last(&u)
	//db.First(&u, "color = ?", "Red")
	// db.Where("color = ?", "Red").
	// 	Where("id > ?", 3).
	// 	First(&u)
	//db.Where(u).First(&u)
	db.Find(&users)
	fmt.Println(len(users))
	fmt.Println(users)
}
