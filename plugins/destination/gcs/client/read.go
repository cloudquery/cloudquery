package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, res chan<- arrow.Record) error {
	if !c.spec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q", table.Name)
	}
	if c.spec.PathContainsUUID() {
		return fmt.Errorf("reading is not supported when path contains uuid variable. Table: %q", table.Name)
	}

	name := c.spec.ReplacePathVariables(table.Name, uuid.NewString(), time.Time{}, c.syncID)

	r, err := c.bucket.Object(name).NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	byteReader := bytes.NewReader(b)
	return c.Client.Read(byteReader, table, res)
}
