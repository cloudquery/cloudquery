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

const (
	batchSize    = 1000
	writeTimeout = 5 * time.Minute
)

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

func (c *Client) writeResource(ctx context.Context, table *schema.Table, client *bigquery.Client, resources <-chan []interface{}) error {
	inserter := client.Dataset(c.pluginSpec.DatasetID).Table(table.Name).Inserter()
	inserter.IgnoreUnknownValues = true
	inserter.SkipInvalidRows = false
	batch := make([]*item, 0)
	for cols := range resources {
		c.logger.Debug().Msg("Got resource")
		saver := &item{
			cols: make(map[string]bigquery.Value, len(table.Columns)),
		}
		for i := range cols {
			if cols[i] == nil {
				// save some bandwidth by not sending nil values
				continue
			}
			saver.cols[table.Columns[i].Name] = cols[i]
		}
		c.logger.Debug().Interface("cols", saver.cols).Msg("got resource")
		batch = append(batch, saver)
		if len(batch) >= batchSize {
			c.logger.Debug().Msg("Writing batch")
			// we use a context with timeout here, because inserter.Put can retry indefinitely
			// on retryable errors if not given a context timeout
			timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
			err := inserter.Put(timeoutCtx, batch)
			if err != nil {
				cancel()
				return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
			}
			// release resources from timeout context if it finished early
			batch = nil
			cancel()
		}
	}
	if len(batch) > 0 {
		c.logger.Debug().Msg("Writing final batch")
		// flush final rows
		timeoutCtx, cancel := context.WithTimeout(ctx, writeTimeout)
		defer cancel()
		err := inserter.Put(timeoutCtx, batch)
		if err != nil {
			return fmt.Errorf("failed to put item into BigQuery table %s: %w", table.Name, err)
		}
	}

	return nil
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *plugins.ClientResource) error {
	eg, gctx := errgroup.WithContext(ctx)
	workers := make(map[string]*worker, len(tables))
	client, err := c.bqClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	for _, t := range tables.FlattenTables() {
		t := t
		writeChan := make(chan []interface{})
		workers[t.Name] = &worker{
			writeChan: writeChan,
		}
		eg.Go(func() error {
			return c.writeResource(gctx, t, client, writeChan)
		})
	}

	done := false
	for !done {
		select {
		case r, ok := <-res:
			if !ok {
				done = true
				break
			}
			workers[r.TableName].writeChan <- r.Data
		case <-gctx.Done():
			done = true
		}
	}
	for _, w := range workers {
		close(w.writeChan)
	}

	return eg.Wait()
}
