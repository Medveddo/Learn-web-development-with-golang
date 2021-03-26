package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "111"
	dbname   = "mywebapp_dev"
)

func main() {
	//psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname) cause an error at db.Ping() pord=%d !!!! ERROR)))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int

	err = db.QueryRow(`
		INSERT INTO  users(name, email)
		VALUES($1,$2)
		RETURNING id`,
		"Kek Ivanovich",
		"kek.lol@gmail.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted user's id: ", id)
	db.Close()
	/*
		C:\Go\src\MyWebApp\exp>go run main.go
		Inserted user's id:  2

		C:\Go\src\MyWebApp\exp>go run main.go
		Inserted user's id:  3

		mywebapp_dev=# SELECT * from users;
		id |      name      |          email
		----+----------------+--------------------------
		1 | Vitaly Sizikov | vitaly.sizikov@gmail.com
		2 | Vitaly Sizikov | vitaly.sizikov@gmail.com
		3 | Kek Ivanovich  | kek.lol@gmail.com
	*/
}
