package client

import (
	"database/sql"
)

func processRows(rows *sql.Rows, process func(row *sql.Rows) error) error {
	defer rows.Close()

	for next := true; next; next = rows.NextResultSet() {
		for rows.Next() {
			if err := process(rows); err != nil {
				return err
			}
		}
	}

	return rows.Err()
}
