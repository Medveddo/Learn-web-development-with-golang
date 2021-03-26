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

	/* INSERTING SOME VALUES
	for i := 1; i <= 6; i++ {
		userID := 1
		if i > 3 {
			userID = 3
		}
		amount := i * 100
		description := fmt.Sprintf("USB-C Adapter x%d", i)

		_, err = db.Exec(`
		INSERT INTO orders(user_id,amount,description)
		VALUES ($1, $2, $3)`, userID, amount, description)

		if err != nil {
			panic(err)
		}
	}
	*/

	rows, err := db.Query(`
	SELECT *
	FROM users
	INNER JOIN orders  ON users.id = orders.user_id`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var userID, orderID, amount int
		var email, name, desc string
		if err := rows.Scan(&userID, &name, &email, &orderID, &userID, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("userID:", userID, "name:", name, "email:", email,
			"orderID:", orderID, "amount:", amount, "desc:", desc)
	}
	if rows.Err() != nil {
		panic(err)
	}
}
