package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	cqtypes "github.com/cloudquery/plugin-sdk/v3/types"
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
		properties[col.Name] = arrowTypeToElasticsearchProperty(col.Type)
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

func arrowTypeToElasticsearchProperty(dataType arrow.DataType) types.Property {
	switch {
	// handle known extensions
	case typeOneOf(dataType,
		cqtypes.ExtensionTypes.UUID,
		cqtypes.ExtensionTypes.MAC,
		cqtypes.ExtensionTypes.Inet):
		return types.NewTextProperty()
	case typeOneOf(dataType,
		cqtypes.ExtensionTypes.JSON):
		return types.NewTextProperty()

	// handle nested types
	case dataType.ID() == arrow.LIST:
		return arrowTypeToElasticsearchProperty(dataType.(*arrow.ListType).Elem())
	case dataType.ID() == arrow.LARGE_LIST:
		return arrowTypeToElasticsearchProperty(dataType.(*arrow.LargeListType).Elem())
	case dataType.ID() == arrow.FIXED_SIZE_LIST:
		return arrowTypeToElasticsearchProperty(dataType.(*arrow.FixedSizeListType).Elem())
	case dataType.ID() == arrow.STRUCT:
		return types.NewObjectProperty()
	case dataType.ID() == arrow.MAP:
		return types.NewObjectProperty()

	// handle primitive types
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Boolean):
		return types.NewBooleanProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Int8):
		return types.NewByteNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Int16):
		return types.NewShortNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Int32):
		return types.NewIntegerNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Int64):
		return types.NewLongNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Uint8,
		arrow.PrimitiveTypes.Uint16,
		arrow.PrimitiveTypes.Uint32,
		arrow.PrimitiveTypes.Uint64):
		return types.NewUnsignedLongNumberProperty()
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Float16):
		return types.NewHalfFloatNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Float32):
		return types.NewFloatNumberProperty()
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Float64):
		return types.NewDoubleNumberProperty()
	case typeOneOf(dataType,
		arrow.BinaryTypes.String,
		arrow.BinaryTypes.LargeString):
		return types.NewTextProperty()
	case typeOneOf(dataType,
		arrow.BinaryTypes.Binary,
		arrow.BinaryTypes.LargeBinary):
		return types.NewBinaryProperty()
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Date32,
		arrow.FixedWidthTypes.Date64):
		d := types.NewDateProperty()
		f := "yyyy-MM-dd"
		d.Format = &f
		return d
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Time32s):
		d := types.NewDateProperty()
		f := "HH:mm:ss"
		d.Format = &f
		return d
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Time32ms):
		d := types.NewDateProperty()
		f := "HH:mm:ss.SSS"
		d.Format = &f
		return d
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Time64us,
		arrow.FixedWidthTypes.Time64ns):
		return types.NewTextProperty()
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Timestamp_s,
		arrow.FixedWidthTypes.Timestamp_ms,
		arrow.FixedWidthTypes.Timestamp_us):
		d := types.NewDateProperty()
		f := "strict_date_optional_time"
		d.Format = &f
		return d
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Timestamp_ns):
		d := types.NewDateNanosProperty()
		f := "strict_date_optional_time_nanos"
		d.Format = &f
	}
	return types.NewTextProperty()
}

func typeOneOf(left arrow.DataType, dt ...arrow.DataType) bool {
	for _, t := range dt {
		if arrow.TypeEqual(left, t) {
			return true
		}
	}
	return false
}
