package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

const maxJsonSize = 1024 * 1024 * 20

func (c *Client) readJSON(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	var rowJson []interface{}
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}
	name := path.Join(c.csvSpec.Directory, table.Name+".csv")
	f, err := c.openReadOnly(ctx, name)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, maxJsonSize), maxJsonSize)
	for scanner.Scan() {
		row := scanner.Bytes()
		if err := json.Unmarshal(row, &rowJson); err != nil {
			return err
		}
		if rowJson[sourceNameIndex] != sourceName {
			continue
		}
		res <- rowJson
	}

	return scanner.Err()
}
