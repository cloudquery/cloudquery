package client

import (
	"context"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when no_rotate is false. Table: %q; Source: %q", tableName, sourceName)
	}
	name := fmt.Sprintf("%s/%s.%s", c.pluginSpec.Directory, tableName, c.pluginSpec.Format)
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.Read(f, arrowSchema, sourceName, res)
}
