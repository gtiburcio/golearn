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

	stmt, _ := db.Prepare("update users set name = ? where id = ?")

	stmt.Exec("Guilherme Tiburcio", 1)

	stmt2, _ := db.Prepare("delete from users where id = ?")

	stmt2.Exec(1)
}
