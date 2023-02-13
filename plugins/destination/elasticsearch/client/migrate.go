package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Migrate creates or updates index templates.
func (c *Client) Migrate(ctx context.Context, tables schema.Tables) error {
	for _, table := range tables {
		tmpl, err := c.getIndexTemplate(table)
		if err != nil {
			return fmt.Errorf("failed to generate index template: %w", err)
		}
		resp, err := c.client.Indices.PutIndexTemplate(table.Name, strings.NewReader(tmpl))
		if err != nil {
			return fmt.Errorf("failed to create index template: %w", err)
		}
		if resp.IsError() {
			return fmt.Errorf("failed to create index template: %s", resp.String())
		}
	}
	return nil
}

func (c *Client) getIndexTemplate(table *schema.Table) (string, error) {
	properties := map[string]types.Property{}
	for _, col := range table.Columns {
		switch col.Type {
		case schema.TypeBool:
			properties[col.Name] = types.NewBooleanProperty()
		case schema.TypeInt:
			properties[col.Name] = types.NewIntegerNumberProperty()
		case schema.TypeFloat:
			properties[col.Name] = types.NewFloatNumberProperty()
		case schema.TypeUUID:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeString:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeByteArray:
			properties[col.Name] = types.NewBinaryProperty()
		case schema.TypeStringArray:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeTimestamp:
			d := types.NewDateProperty()
			f := "strict_date_optional_time||epoch_millis"
			d.Format = &f
			properties[col.Name] = d
		case schema.TypeJSON:
			properties[col.Name] = types.NewObjectProperty()
		case schema.TypeUUIDArray:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeCIDR:
			properties[col.Name] = types.NewIpRangeProperty()
		case schema.TypeCIDRArray:
			properties[col.Name] = types.NewIpRangeProperty()
		case schema.TypeMacAddr:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeMacAddrArray:
			properties[col.Name] = types.NewTextProperty()
		case schema.TypeInet:
			properties[col.Name] = types.NewIpRangeProperty()
		case schema.TypeInetArray:
			properties[col.Name] = types.NewIpRangeProperty()
		case schema.TypeIntArray:
			properties[col.Name] = types.NewIntegerNumberProperty()
		}
	}
	tmp := types.IndexTemplate{
		AllowAutoCreate: nil,
		ComposedOf:      []string{},
		DataStream:      nil,
		IndexPatterns:   []string{c.getIndexNamePattern(table.Name)},
		Meta_:           nil,
		Priority:        nil,
		Template: &types.IndexTemplateSummary{
			Mappings: &types.TypeMapping{
				Properties: properties,
			},
		},
		Version: nil,
	}
	b, err := json.Marshal(tmp)
	return string(b), err
}
