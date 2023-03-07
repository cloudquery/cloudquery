package main

import (
	"database/sql"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE foo (id int PRIMARY KEY, s text);")
	if err != nil {
		panic(err)
	}
	s := "AStringWith\u0000NullBytes"
	_, err = db.Exec("INSERT INTO foo VALUES (?, ?);", 1, s)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT s FROM foo;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var got string
		rows.Scan(&got)
		if got != s {
			fmt.Println("got", got, "want", s)
		} else {
			fmt.Println("PASS")
		}
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
