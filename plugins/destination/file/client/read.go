package client

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) readLocal(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	filePath := path.Join(c.csvSpec.Directory, table.Name+".csv")
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip header
	_, err = r.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}

	for {
		record, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if record[sourceNameIndex] != sourceName {
			continue
		}
		values := make([]interface{}, len(record))
		for i, v := range record {
			values[i] = v
		}

		res <- values
	}
	return nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	switch c.csvSpec.Backend {
	case BackendTypeLocal:
		return c.readLocal(ctx, table, sourceName, res)
	default:
		panic("unknown backend type " + c.csvSpec.Backend)
	}
}
