package client

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) read(table *schema.Table, sourceName string, res chan<- *schema.DestinationResource) error {
	filePath := path.Join(c.csvSpec.Directory, table.Name+".csv")
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip header
	_, err = r.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}
	sourceNameIndex := table.Columns.Index(schema.CqSourceNameColumn.Name)
	if sourceNameIndex == -1 {
		return fmt.Errorf("could not find column %s in table %s", schema.CqSourceNameColumn.Name, table.Name)
	}

	for {
		record, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if record[sourceNameIndex] != sourceName {
			continue
		}

		cqTypes, err := schema.CQTypesFromValues(table, record)
		if err != nil {
			return err
		}
		res <- &schema.DestinationResource{
			TableName: table.Name,
			Data:      cqTypes,
		}
	}
	return nil
}

func (c *Client) Read(tx context.Context, table *schema.Table, sourceName string, res chan<- *schema.DestinationResource) error {
	msg := &readMsg{
		table:     table,
		source:    sourceName,
		err:       make(chan error),
		resources: make(chan *schema.DestinationResource),
	}
	c.readChan <- msg
	for {
		select {
		case err := <-msg.err:
			return err
		case r := <-msg.resources:
			res <- r
		}
	}
}
