package client

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, sc *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(sc)
	f, err := os.CreateTemp("", fmt.Sprintf("%s-*.json", tableName))
	if err != nil {
		return err
	}
	fName := f.Name()
	if err := f.Close(); err != nil {
		return err
	}

	_, err = c.db.Exec("copy " + tableName + " to '" + f.Name() + "' (timestampformat '%Y-%m-%d %H:%M:%S.%f')")
	if err != nil {
		return err
	}
	f, err = os.Open(fName)
	if err != nil {
		return err
	}

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(f)

	// Loop through the scanner, reading line by line
	for scanner.Scan() {
		line := scanner.Bytes()
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, sc)
		if err := bldr.UnmarshalJSON(line); err != nil {
			return err
		}
		res <- bldr.NewRecord()
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading temporary json file: %s", err)
	}

	return nil
}
