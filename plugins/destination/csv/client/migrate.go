package client

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

func isSameColumns(record []string, columns []string) bool {
	if len(record) != len(columns) {
		return false
	}
	for i, v := range record {
		if v != columns[i] {
			return false
		}
	}
	return true
}

func (c *Client) migrate(tables schema.Tables) error {
	for _, t := range tables {
		filePath := path.Join(c.csvSpec.Directory, t.Name+".csv")
		//nolint:gocritic,revive
		if _, err := os.Stat(filePath); err == nil {
			if err := func() error {
				f, err := os.Open(filePath)
				if err != nil {
					return err
				}
				defer f.Close()
				r := csv.NewReader(f)
				// skip header
				record, err := r.Read()
				if err != nil {
					return err
				}
				if !isSameColumns(record, t.Columns.Names()) {
					return fmt.Errorf("csv can't migrate table %s. please delete %s file and try again", t.Name, filePath)
				}
				return nil
			}(); err != nil {
				return err
			}
		} else if errors.Is(err, os.ErrNotExist) {
			f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				return err
			}
			w := csv.NewWriter(f)
			if err := w.Write(t.Columns.Names()); err != nil {
				f.Close()
				return err
			}
			w.Flush()
			if err := f.Close(); err != nil {
				return err
			}
		} else {
			return err
		}

		if err := c.migrate(t.Relations); err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	msg := &migrateMsg{
		tables: tables,
		err:    make(chan error),
	}
	c.migrateChan <- msg
	return <-msg.err
}
