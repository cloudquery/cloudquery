package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

func (c *Client) ColumnToBigQuerySchema(col schema.Column) *bigquery.FieldSchema {
	sc := bigquery.FieldSchema{
		Name:        col.Name,
		Description: col.Description,
		// Repeated:               false,
		Required: col.NotNull,
		// Type:                   "",
		// Schema:                 nil,
		// MaxLength:              0,
		// Precision:              0,
		// Scale:                  0,
		// DefaultValueExpression: "",
		// Collation:              "",
	}

	switch v := col.Type.(type) {
	case *arrow.ListType:
		ift, sch, isRepeated := c.ArrowTypeToBigQuery(v.Elem())
		if isRepeated {
			// store nested arrays as JSON
			return bigquery.JSONFieldType, nil, false
		}
		return ift, sch, true
	case *arrow.MapType:
		// store maps as JSON
		return bigquery.JSONFieldType, nil, false
	}

	switch {
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Boolean):
		sc.Type = bigquery.BooleanFieldType
	case typeOneOf(col.Type,
		arrow.PrimitiveTypes.Int8,
		arrow.PrimitiveTypes.Int16,
		arrow.PrimitiveTypes.Int32,
		arrow.PrimitiveTypes.Int64,
		arrow.PrimitiveTypes.Uint8,
		arrow.PrimitiveTypes.Uint16,
		arrow.PrimitiveTypes.Uint32,
		arrow.PrimitiveTypes.Uint64):
		sc.Type = bigquery.IntegerFieldType
	case typeOneOf(col.Type,
		arrow.PrimitiveTypes.Float32,
		arrow.PrimitiveTypes.Float64):
		sc.Type = bigquery.FloatFieldType
	case typeOneOf(col.Type,
		arrow.BinaryTypes.String,
		arrow.BinaryTypes.LargeString):
		sc.Type = bigquery.StringFieldType
	case typeOneOf(col.Type,
		arrow.BinaryTypes.Binary,
		arrow.BinaryTypes.LargeBinary):
		sc.Type = bigquery.BytesFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Date32,
		arrow.FixedWidthTypes.Date64):
		sc.Type = bigquery.DateFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Timestamp_s,
		arrow.FixedWidthTypes.Timestamp_ms,
		arrow.FixedWidthTypes.Timestamp_us,
		arrow.FixedWidthTypes.Timestamp_ns):
		sc.Type = bigquery.TimestampFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Time32s,
		arrow.FixedWidthTypes.Time64us):
		sc.Type = bigquery.TimeFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Time64ns):
		// we lose nanosecond precision here. BigQuery doesn't support it
		// so we need to live with this for now.
		sc.Type = bigquery.TimeFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.Duration_s,
		arrow.FixedWidthTypes.Duration_ms,
		arrow.FixedWidthTypes.Duration_us,
		arrow.FixedWidthTypes.Duration_ns):
		sc.Type = bigquery.IntervalFieldType
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.MonthInterval):
		sc.Type = bigquery.RecordFieldType
		sc.Schema = []*bigquery.FieldSchema{
			{
				Name: "months",
				Type: bigquery.IntegerFieldType,
			},
		}
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.DayTimeInterval):
		sc.Type = bigquery.RecordFieldType
		sc.Schema = []*bigquery.FieldSchema{
			{
				Name: "days",
				Type: bigquery.IntegerFieldType,
			},
			{
				Name: "milliseconds",
				Type: bigquery.IntegerFieldType,
			},
		}
	case typeOneOf(col.Type,
		arrow.FixedWidthTypes.MonthDayNanoInterval):
		sc.Type = bigquery.RecordFieldType
		sc.Schema = []*bigquery.FieldSchema{
			{
				Name: "months",
				Type: bigquery.IntegerFieldType,
			},
			{
				Name: "days",
				Type: bigquery.IntegerFieldType,
			},
			{
				Name: "nanoseconds",
				Type: bigquery.IntegerFieldType,
			},
		}
	case typeOneOf(col.Type,
		types.ExtensionTypes.Inet,
		types.ExtensionTypes.MAC,
		types.ExtensionTypes.UUID):
		sc.Type = bigquery.StringFieldType
	case typeOneOf(col.Type,
		types.ExtensionTypes.JSON):
		sc.Type = bigquery.JSONFieldType
	}

	return &sc
}

func typeOneOf(left arrow.DataType, dt ...arrow.DataType) bool {
	for _, t := range dt {
		if arrow.TypeEqual(left, t) {
			return true
		}
	}
	return false
}
