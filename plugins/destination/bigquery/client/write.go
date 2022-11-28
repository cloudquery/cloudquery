package client

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

const batchSize = 1000

type worker struct {
	writeChan chan []interface{}
}

type item struct {
	cols map[string]bigquery.Value
}

func (i *item) Save() (map[string]bigquery.Value, string, error) {
	// we're not doing de-dup at the moment
	return i.cols, bigquery.NoDedupeID, nil
}

func (c *Client) writeResource(ctx context.Context, table *schema.Table, resources <-chan []interface{}) error {
	inserter := c.client.Dataset(c.datasetID).Table(table.Name).Inserter()
	batch := make([]item, 0)
	for r := range resources {
		c.logger.Debug().Msg("Got resource")
		saver := item{
			cols: make(map[string]bigquery.Value, len(table.Columns)),
		}
		for i := range r {
			saver.cols[table.Columns[i].Name] = r[i]
		}
		c.logger.Debug().Interface("cols", saver.cols).Msg("got resource")
		batch = append(batch, saver)
		if len(batch) >= batchSize {
			c.logger.Debug().Msg("Writing batch")
			// we use a context with timeout here, because inserter.Put can retry indefinitely
			// on retryable errors if not given a context timeout
			timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			err := inserter.Put(timeoutCtx, batch)
			if err != nil {
				return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
			}
			// release resources from timeout context if it finished early
			cancel()
		}
	}
	if len(batch) > 0 {
		c.logger.Debug().Msg("Writing final batch")
		// flush final rows
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		err := inserter.Put(timeoutCtx, batch)
		if err != nil {
			return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
		}
		cancel()
	}

	return nil
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	eg := errgroup.Group{}
	workers := make(map[string]*worker, len(tables))

	for _, t := range tables.FlattenTables() {
		t := t
		writeChan := make(chan []interface{})
		workers[t.Name] = &worker{
			writeChan: writeChan,
		}
		eg.Go(func() error {
			return c.writeResource(ctx, t, writeChan)
		})
	}

	for r := range res {
		workers[r.TableName].writeChan <- r.Data
	}
	for _, w := range workers {
		close(w.writeChan)
	}

	return eg.Wait()
}
