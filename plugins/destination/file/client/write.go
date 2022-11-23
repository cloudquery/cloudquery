package client

import (
	"context"
	"sync"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

type worker struct {
	writeChan chan []interface{}
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	wg := sync.WaitGroup{}
	workers := make(map[string]*worker, len(tables))
	for _, t := range tables.FlattenTables() {
		t := t
		workers[t.Name] = &worker{
			writeChan: make(chan []interface{}),
		}
		wg.Add(1)
		switch c.csvSpec.Format {
		case FormatTypeCSV:
			go func() {
				defer wg.Done()
				c.writeCSVResource(ctx, t.Name, workers[t.Name].writeChan)
			}()
		case FormatTypeJSON:
			go func() {
				defer wg.Done()
				c.writeJSONResource(ctx, t, workers[t.Name].writeChan)
			}()
		default:
			panic("unknown format type")
		}
	}

	for r := range res {
		workers[r.TableName].writeChan <- r.Data
	}
	for _, w := range workers {
		close(w.writeChan)
	}

	wg.Wait()
	return nil
}
