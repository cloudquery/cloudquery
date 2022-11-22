package client

import (
	"context"
	"encoding/json"
	"path"
)

func (c *Client) writeJSONResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	filePath := path.Join(c.csvSpec.Directory, tableName+".json")
	var err error
	f, err := c.OpenAppendOnly(ctx, filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	for r := range resources {
		b, err := json.Marshal(r)
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
