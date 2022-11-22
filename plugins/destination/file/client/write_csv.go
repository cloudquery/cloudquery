package client

import (
	"context"
	"encoding/csv"
	"path"
)

func (c *Client) writeCSVResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	var err error
	filePath := path.Join(c.csvSpec.Directory, tableName+".csv")
	f, err := c.OpenAppendOnly(ctx, filePath)
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
