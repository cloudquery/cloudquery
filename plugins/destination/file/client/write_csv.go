package client

import (
	"context"
	"encoding/csv"
	"os"
	"path"
)

func (c *Client) writeCSVResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	var err error
	var records [][]string
	var f *os.File
	totalRecords := uint64(0)
	filePath := path.Join(c.csvSpec.Directory, tableName+".csv")
	if c.csvSpec.Backend == BackendTypeLocal {
		f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	for r := range resources {
		record := make([]string, len(r))
		for i, v := range r {
			record[i] = v.(string)
		}
		records = append(records, record)
		totalRecords++
		if totalRecords >= c.csvSpec.BatchSize {
			csvWriter := csv.NewWriter(f)
			if err := csvWriter.WriteAll(records); err != nil {
				return err
			}
			records = nil
			totalRecords = 0
		}
	}
	if totalRecords > 0 {
		csvWriter := csv.NewWriter(f)
		if err := csvWriter.WriteAll(records); err != nil {
			return err
		}
	}
	return nil
}
