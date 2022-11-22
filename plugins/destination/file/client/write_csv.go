package client

import (
	"context"
	"encoding/csv"
	"os"
	"path"
)

func (c *Client) writeCSVResource(_ context.Context, tableName string, resources <-chan []interface{}) error {
	var err error
	var f *os.File
	filePath := path.Join(c.csvSpec.Directory, tableName+".csv")
	if c.csvSpec.Backend == BackendTypeLocal {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
	}
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
