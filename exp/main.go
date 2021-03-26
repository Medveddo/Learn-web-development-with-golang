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
	var name, email string

	row := db.QueryRow(`
	SELECT id, name, email
	FROM users
	WHERE id=$1`, 3)
	err = row.Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			panic(err)
		}
	}
	fmt.Println("id", id, "name", name, "email", email)
}
