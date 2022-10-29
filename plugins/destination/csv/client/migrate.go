package client

import (
	"context"
	"encoding/csv"
	"errors"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) migrate(tables schema.Tables) error {
	for _, t := range tables {
		filePath := path.Join(c.csvSpec.Directory, t.Name+".csv")
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
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
