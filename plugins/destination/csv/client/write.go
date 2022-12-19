package client

import (
	"context"
	"encoding/csv"
	"os"
	"path"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) releaseTables(tables []string) {
	for _, t := range tables {
		if c.writers[t].count == 0 {
			panic("unexpected count is 0")
		}
		c.writers[t].count--
		if c.writers[t].count == 0 {
			c.writers[t].writer.Flush()
			fileName := c.writers[t].file.Name()
			if err := c.writers[t].file.Close(); err != nil {
				c.logger.Error().Err(err).Msgf("failed to close file %s", fileName)
			}
			delete(c.writers, t)
		}
	}
}

func (c *Client) startWrite(tables schema.Tables) error {
	// we want to have the ability to rollback in case of an error so we dont have zombie files
	var initializedTables []string
	for _, t := range tables {
		if c.writers[t.Name] == nil {
			filePath := path.Join(c.csvSpec.Directory, t.Name+".csv")
			f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				c.logger.Error().Err(err).Str("table", t.Name).Msgf("failed to open file %s", filePath)
				c.metrics.Errors++
				c.releaseTables(initializedTables)
				return err
			}
			c.writers[t.Name] = &tableWriter{
				writer: csv.NewWriter(f),
				file:   f,
				count:  1,
			}
			initializedTables = append(initializedTables, t.Name)
		} else {
			c.writers[t.Name].count++
		}
		if err := c.startWrite(t.Relations); err != nil {
			c.releaseTables(initializedTables)
			return err
		}
	}
	return nil
}

func (c *Client) endWrite(tables schema.Tables) {
	for _, t := range tables {
		if c.writers[t.Name] != nil {
			c.writers[t.Name].writer.Flush()
			if c.writers[t.Name].count == 0 {
				panic("unexpected count is 0")
			}
			c.writers[t.Name].count--
			if c.writers[t.Name].count == 0 {
				fileName := c.writers[t.Name].file.Name()
				if err := c.writers[t.Name].file.Close(); err != nil {
					c.logger.Error().Err(err).Msgf("failed to close file %s", fileName)
				}
				delete(c.writers, t.Name)
			}
		}
		c.endWrite(t.Relations)
	}
}

// this should only be called from the main listen goroutine to ensure writing to files
// only happens in one place
func (c *Client) writeResource(resource *destination.ClientResource) {
	record := make([]string, len(resource.Data))
	for i, v := range resource.Data {
		record[i] = v.(string)
	}

	if err := c.writers[resource.TableName].writer.Write(record); err != nil {
		c.metrics.Errors++
		c.logger.Error().Err(err).Str("table", resource.TableName).Msg("failed to write resource")
	}
}

func (c *Client) Write(ctx context.Context, tables schema.Tables, res <-chan *destination.ClientResource) error {
	startWriteMsg := &startWriteMsg{
		tables: tables,
		err:    make(chan error),
	}
	c.startWriteChan <- startWriteMsg
	if err := <-startWriteMsg.err; err != nil {
		return err
	}

	for r := range res {
		c.writeChan <- r
	}

	endWriteMsg := &endWriteMsg{
		tables: tables,
		err:    make(chan error),
	}
	c.endWriteChan <- endWriteMsg
	return <-endWriteMsg.err
}
