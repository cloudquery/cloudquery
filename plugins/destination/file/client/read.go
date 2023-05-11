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

func (c *Client) Read(ctx context.Context, arrowSchema *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	tableName := schema.TableName(arrowSchema)
	if !c.pluginSpec.NoRotate {
		return fmt.Errorf("reading is not supported when `no_rotate` is false. Table: %q; Source: %q", tableName, sourceName)
	}
	if strings.Contains(c.pluginSpec.Path, PathVarUUID) {
		return fmt.Errorf("reading is not supported when `path` contains UUID variable. Table: %q; Source: %q", tableName, sourceName)
	}
	name := replacePathVariables(c.pluginSpec.Path, tableName, c.pluginSpec.Format, uuid.NewString(), time.Time{})
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return c.Client.Read(f, arrowSchema, sourceName, res)
}
