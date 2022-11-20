package client

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path"
)

func (c *Client) writeJSONResource(ctx context.Context, tableName string, resources <-chan []interface{}) error {
	var buf bytes.Buffer
	totalRecords := uint64(0)
	filePath := path.Join(c.csvSpec.Directory, tableName+".csv")
	var f *os.File
	if c.csvSpec.Backend == BackendTypeLocal {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	for r := range resources {
		b, err := json.Marshal(r)
		if err != nil {
			return err
		}
		_, err = buf.Write(b)
		if err != nil {
			return err
		}
		_, err = buf.Write([]byte("\n"))
		if err != nil {
			return err
		}
		totalRecords++
		if totalRecords >= c.csvSpec.BatchSize {
			if c.csvSpec.Backend == BackendTypeLocal {
				_, err = f.Write(buf.Bytes())
				if err != nil {
					return err
				}
				if err := f.Sync(); err != nil {
					return err
				}
			} else {
				if err := c.writeRemote(ctx, tableName, buf.Bytes()); err != nil {
					return err
				}
			}
		}
		buf = bytes.Buffer{}
		totalRecords = 0
	}

	return nil
}
