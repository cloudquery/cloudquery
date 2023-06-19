package client

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
)

func (c *Client) Read(_ context.Context, table *schema.Table, res chan<- arrow.Record) error {
	sourceName := "unknown"

	if !c.spec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	if strings.Contains(c.spec.Path, PathVarUUID) {
		return fmt.Errorf("reading is not supported when `path` contains UUID variable. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := replacePathVariables(c.spec.Path, table.Name, c.spec.Format, uuid.NewString(), time.Time{})
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.Read(f, table, sourceName, res)
}
