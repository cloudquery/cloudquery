package client

import (
	"context"
	"encoding/csv"
)

func (c *Client) writeCSVResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	var err error
	f, err := c.openAppendOnly(ctx, tableName+".csv")
	if err != nil {
		return err
	}
	defer f.Close()
	csvWriter := csv.NewWriter(f)

	for r := range resources {
		record := make([]string, len(r))
		for i, v := range r {
			record[i] = v.(string)
		}
		csvWriter.Write(record)
	}
	csvWriter.Flush()
	return nil
}
