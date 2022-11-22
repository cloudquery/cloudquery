package client

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) readCSV(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	name := path.Join(c.csvSpec.Directory, table.Name+".csv")
	reader, err := c.openReadOnly(ctx, name)
	if err != nil {
		return err
	}
	defer reader.Close()

	r := csv.NewReader(reader)
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
