package main

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE foo (b uuid);")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO foo VALUES ('eeccb8c5-9943-b2bb-bb5e-222f4e14b687');")
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
