package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userById(w, r, id)
	case r.Method == "GET":
		findAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found.")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/gocurse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var user User

	db.QueryRow("select id, name from users where id = ?", id)

	json, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/gocurse")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, name from users")

	var users []User

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
