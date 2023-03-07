package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/marcboeker/go-duckdb"
	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE foo (id int PRIMARY KEY, b text, u uuid[]);")
	if err != nil {
		panic(err)
	}
	//myList := []string{"a", "b"}
	_, err = db.Exec("INSERT INTO foo VALUES (?, ?, [?, ?]);", 1, "a", uuid.New().String(), uuid.New().String())
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT u FROM foo;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var l duckdb.Composite[[][]byte]
		rows.Scan(&l)
		fmt.Println(uuid.FromBytes(l.Get()[0]))
	}
	//
	//_, err = db.Exec("INSERT OR REPLACE INTO foo VALUES (?, ?);", 1, "cb")
	//if err != nil {
	//	panic(err)
	//}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
