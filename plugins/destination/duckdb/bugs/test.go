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
	_, err = db.Exec("CREATE TABLE foo (b text[]);")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO foo VALUES ($1);", `['a', 'b']`)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT b FROM foo;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var l string
		rows.Scan(&l)
		fmt.Println(l)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
}
