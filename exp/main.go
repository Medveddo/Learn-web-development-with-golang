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

	type User struct {
		ID    int
		Name  string
		Email string
	}
	var users []User

	rows, err := db.Query(`
	SELECT id, name, email
	FROM users`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		// handle this err
	}
	fmt.Println(users)
}
