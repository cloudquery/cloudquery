package client

import (
	"container/heap"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/cloudquery/cloudquery/plugins/destination/csv/internal/priorityqueue"
	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) read(table *schema.Table, sourceName string, res chan<- []any, opts destination.ReadOptions) error {
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

	// priority queue is used when order by is specified. We use it to sort the rows in memory
	// without reading the entire file into memory (as long as a limit is applied).
	pq := priorityqueue.New(table, opts.OrderBy)
	for count := 0; opts.Limit == 0 || count < opts.Limit; {
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

		var values []any
		if len(opts.Columns) == 0 || len(opts.OrderBy) > 0 {
			// if order by is set, we need to read all columns for the priority
			// queue to be able to index columns correctly
			values = make([]any, len(record))
			for i, v := range record {
				values[i] = v
			}
		} else {
			values = make([]any, len(opts.Columns))
			for i, col := range opts.Columns {
				ind := table.Columns.Index(col)
				if ind == -1 {
					return fmt.Errorf("could not find column %s in table %s", col, table.Name)
				}
				values[i] = record[ind]
			}
		}

		if len(opts.OrderBy) == 0 {
			res <- values
			count++
		} else {
			t, err := c.ReverseTransformValues(table, values)
			if err != nil {
				return fmt.Errorf("failed to reverse transform values: %w", err)
			}
			heap.Push(pq, priorityqueue.NewItem(t))
			if opts.Limit > 0 && pq.Len() > opts.Limit {
				heap.Pop(pq)
			}
		}
	}
	if len(opts.OrderBy) > 0 {
		n := pq.Len()
		final := make([][]any, n)
		for i := 0; i < n; i++ {
			it := pq.Pop().(*priorityqueue.Item)
			cols := make([]any, len(it.Cols))
			for vi, v := range it.Cols {
				cols[vi] = v.Get()
			}
			final[i] = cols
		}
		for _, v := range final {
			res <- v
		}
	}
	return nil
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any, opts destination.ReadOptions) error {
	msg := &readMsg{
		table:     table,
		source:    sourceName,
		options:   opts,
		err:       make(chan error),
		resources: make(chan []any),
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
