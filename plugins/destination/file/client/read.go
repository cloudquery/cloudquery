package client

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/google/uuid"
)

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- arrow.Record) error {
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", table.Name, sourceName)
	}
	if strings.Contains(c.pluginSpec.Path, PathVarUUID) {
		return fmt.Errorf("reading is not supported when `path` contains UUID variable. Table: %q; Source: %q", table.Name, sourceName)
	}
	name := replacePathVariables(c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString(), time.Time{})
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.Read(f, table, sourceName, res)
}
