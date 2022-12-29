package client

import (
	"context"
	"fmt"

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
	name := fmt.Sprintf("%s/%s.%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format, uuid.NewString())
	if c.pluginSpec.NoRotate {
		name = fmt.Sprintf("%s/%s.%s", c.pluginSpec.Path, table.Name, c.pluginSpec.Format)
	}
	r, err := c.bucket.Object(name).NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	switch c.pluginSpec.Format {
	case FormatTypeCSV:
		if err := csv.Read(r, table, sourceName, res); err != nil {
			return err
		}
	case FormatTypeJSON:
		if err := json.Read(r, table, sourceName, res); err != nil {
			return err
		}
	default:
		panic("unknown format " + c.pluginSpec.Format)
	}
	return nil
}
