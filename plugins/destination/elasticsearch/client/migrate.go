package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"golang.org/x/exp/maps"
)

// Migrate creates or updates index templates.
func (c *Client) MigrateTables(ctx context.Context, msgs message.WriteMigrateTables) error {
	for _, msg := range msgs {
		table := msg.Table
		tmpl, err := c.getIndexTemplate(table)
		if err != nil {
			return fmt.Errorf("failed to generate index template: %w", err)
		}

		if msg.MigrateForce {
			var indicesToDelete []string
			pat := c.getIndexNamePattern(table)
			if strings.HasSuffix(pat, "*") {
				resp, err := c.client.Indices.Get([]string{pat},
					c.client.Indices.Get.WithContext(ctx),
					c.client.Indices.Get.WithIgnoreUnavailable(true),
					c.client.Indices.Get.WithFeatures("aliases"),
				)
				if err != nil {
					return fmt.Errorf("failed to get indices: %w", err)
				}
				if resp.IsError() {
					return fmt.Errorf("failed to get indices: %s", resp.String())
				}

				var indices map[string]any
				if err := json.NewDecoder(resp.Body).Decode(&indices); err != nil {
					return fmt.Errorf("failed to decode response body: %w", err)
				}
				_ = resp.Body.Close()

				indicesToDelete = maps.Keys(indices)
			} else {
				indicesToDelete = []string{table.Name}
			}

			if len(indicesToDelete) > 0 {
				if err := c.deleteIndices(ctx, indicesToDelete); err != nil {
					return fmt.Errorf("failed to delete indices: %w", err)
				}
			}
		}

		resp, err := c.client.Indices.PutIndexTemplate(
			table.Name,
			strings.NewReader(tmpl),
			c.client.Indices.PutIndexTemplate.WithContext(ctx),
			c.client.Indices.PutIndexTemplate.WithCreate(false),
		)
		if err != nil {
			return fmt.Errorf("failed to create index template: %w", err)
		}
		if resp.IsError() {
			return fmt.Errorf("failed to create index template: %s", resp.String())
		}

		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
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
		IndexPatterns:   []string{c.getIndexNamePattern(table)},
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

func (c *Client) deleteIndices(ctx context.Context, names []string) error {
	c.logger.Debug().Strs("indices", names).Msg("deleting indices")
	resp, err := c.client.Indices.Delete(names,
		c.client.Indices.Delete.WithContext(ctx),
		c.client.Indices.Delete.WithIgnoreUnavailable(true),
	)
	if err != nil {
		return fmt.Errorf("failed to delete indices: %w", err)
	}
	defer resp.Body.Close()

	if resp.IsError() {
		return fmt.Errorf("failed to delete indices: %s", resp.String())
	}
	_, _ = io.Copy(io.Discard, resp.Body)

	return nil
}

func arrowTypeToElasticsearchProperty(dataType arrow.DataType) types.Property {
	if dataType == nil {
		return types.NewTextProperty()
	}
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
		p := types.NewObjectProperty()
		for _, field := range dataType.(*arrow.StructType).Fields() {
			p.Properties[field.Name] = arrowTypeToElasticsearchProperty(field.Type)
		}
		return p
	case dataType.ID() == arrow.MAP:
		p := types.NewObjectProperty()
		p.Properties["key"] = arrowTypeToElasticsearchProperty(dataType.(*arrow.MapType).KeyType())
		p.Properties["value"] = arrowTypeToElasticsearchProperty(dataType.(*arrow.MapType).ItemType())
		return p
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
		return d
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.DayTimeInterval,
		arrow.FixedWidthTypes.MonthInterval,
		arrow.FixedWidthTypes.MonthDayNanoInterval):
		return types.NewObjectProperty()
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
