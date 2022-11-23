package client

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/schema"
)

const maxJsonSize = 1024 * 1024 * 20

func (c *Client) writeJSONResource(ctx context.Context, table *schema.Table, resources <-chan []interface{}) error {
	f, err := c.openAppendOnly(ctx, table.Name+".json")
	if err != nil {
		return err
	}
	defer f.Close()


	for r := range resources {
		jsonObj := make(map[string]interface{}, len(table.Columns))
		for i := range r {
			jsonObj[table.Columns[i].Name] = r[i]
		}

		b, err := json.Marshal(jsonObj)
		if err != nil {
			return err
		}
		b = append(b, '\n')
		if _, err := f.Write(b); err != nil {
			return err
		}
	}

	return err
}

func (c *Client) readJSON(ctx context.Context, table *schema.Table, sourceName string, res chan<- []interface{}) error {
	var rowJson []interface{}
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}
	name := table.Name + ".json"
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
