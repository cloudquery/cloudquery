package client

import (
	"context"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
)

func (c *Client) read(tables schema.Tables, sourceName string, syncTime time.Time) error {
	return nil
}

func (c *Client) Read(ctx context.Context, tables schema.Tables, sourceName string, syncTime time.Time) error {
	msg := &readMsg{
		tables: tables,
		source: sourceName,
		sync:   syncTime,
		err:    make(chan error),
	}
	c.readChan <- msg
	return <-msg.err
}