package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, err := sql.Open("mysql", "root:root@/gocurse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id > ?", 2)
	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		users = append(users, u)
	}

	for _, u := range users {
		fmt.Println(u)
	}
}
