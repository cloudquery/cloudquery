package client

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/filetypes/csv"
	"github.com/cloudquery/filetypes/json"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/uuid"
)

func (c *Client) ReverseTransformValues(table *schema.Table, values []any) (schema.CQTypes, error) {
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		return c.csvReverseTransformer.ReverseTransformValues(table, values)
	case FormatTypeJSON:
		return c.jsonReverseTransformer.ReverseTransformValues(table, values)
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
}

func (c *Client) Read(ctx context.Context, table *schema.Table, sourceName string, res chan<- []any) error {
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Directory, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Directory, table.Name, c.pluginSpec.Format)
	}
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		if err := csv.Read(f, table, sourceName, res); err != nil {
			return err
		}
	case FormatTypeJSON:
		if err := json.Read(f, table, sourceName, res); err != nil {
			return err
		}
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
	return nil
}
