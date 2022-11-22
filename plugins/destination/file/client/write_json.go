package client

import (
	"context"
	"encoding/json"
)

func (c *Client) writeJSONResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	f, err := c.openAppendOnly(ctx, tableName+".json")
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
