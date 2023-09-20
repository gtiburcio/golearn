package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/gocurse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into users(id, name) values (?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, err = stmt.Exec(2001, "Tiago")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
