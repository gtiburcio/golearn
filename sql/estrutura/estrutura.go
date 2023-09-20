package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists gocurse")
	exec(db, "use gocurse")
	exec(db, `create table if not exists users (
		id integer auto_increment,
		name varchar(90),
		PRIMARY KEY (id)
	)`)
}
