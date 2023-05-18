package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

// TimestampNanoseconds is a struct to hold a timestamp with nanosecond precision,
// because BigQuery does not support nanosecond precision timestamps.
type TimestampNanoseconds struct {
	Timestamp   string `json:"timestamp" bigquery:"timestamp"`
	Nanoseconds int    `json:"nanoseconds" bigquery:"nanoseconds"`
}

// DurationNanoseconds is a struct to hold a duration with nanosecond precision,
// because BigQuery does not support nanosecond precision durations.
type DurationNanoseconds struct {
	Duration    string `json:"duration" bigquery:"duration"`
	Nanoseconds int    `json:"nanoseconds" bigquery:"nanoseconds"`
}

func (c *Client) ColumnToBigQuerySchema(col schema.Column) *bigquery.FieldSchema {
	sc := bigquery.FieldSchema{
		Name:        col.Name,
		Description: col.Description,
		// Required:    col.NotNull,
		Type:   c.DataTypeToBigQueryType(col.Type),
		Schema: c.DataTypeToBigQuerySchema(col.Type),
	}

	if col.Type.ID() == arrow.LIST {
		sc.Repeated = true
	}
	return &sc
}

func (c *Client) DataTypeToBigQueryType(dataType arrow.DataType) bigquery.FieldType {
	// handle complex types
	switch dataType.ID() {
	case arrow.MAP:
		return bigquery.JSONFieldType
	case arrow.STRUCT:
		return bigquery.RecordFieldType
	case arrow.LIST:
		switch v := dataType.(type) {
		case *arrow.ListType:
			return c.DataTypeToBigQueryType(v.Elem())
		case *arrow.LargeListType:
			return c.DataTypeToBigQueryType(v.Elem())
		}
	}

	// handle basic types
	switch {
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Boolean):
		return bigquery.BooleanFieldType
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Int8,
		arrow.PrimitiveTypes.Int16,
		arrow.PrimitiveTypes.Int32,
		arrow.PrimitiveTypes.Int64,
		arrow.PrimitiveTypes.Uint8,
		arrow.PrimitiveTypes.Uint16,
		arrow.PrimitiveTypes.Uint32,
		arrow.PrimitiveTypes.Uint64):
		return bigquery.IntegerFieldType
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Float32,
		arrow.PrimitiveTypes.Float64):
		return bigquery.FloatFieldType
	case typeOneOf(dataType,
		arrow.BinaryTypes.String,
		arrow.BinaryTypes.LargeString):
		return bigquery.StringFieldType
	case typeOneOf(dataType,
		arrow.BinaryTypes.Binary,
		arrow.BinaryTypes.LargeBinary):
		return bigquery.BytesFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Date32,
		arrow.FixedWidthTypes.Date64):
		return bigquery.DateFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Timestamp_s,
		arrow.FixedWidthTypes.Timestamp_ms,
		arrow.FixedWidthTypes.Timestamp_us):
		return bigquery.TimestampFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Timestamp_ns):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Time32s,
		arrow.FixedWidthTypes.Time32ms,
		arrow.FixedWidthTypes.Time64us,
		arrow.FixedWidthTypes.Time64ns):
		return bigquery.TimeFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Duration_s,
		arrow.FixedWidthTypes.Duration_ms,
		arrow.FixedWidthTypes.Duration_us):
		return bigquery.IntervalFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.Duration_ns):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.MonthInterval):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.DayTimeInterval):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.MonthDayNanoInterval):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		types.ExtensionTypes.Inet,
		types.ExtensionTypes.MAC,
		types.ExtensionTypes.UUID):
		return bigquery.StringFieldType
	case typeOneOf(dataType,
		types.ExtensionTypes.JSON):
		return bigquery.JSONFieldType
	default:
		panic("unsupported data type: " + dataType.String())
	}
}

func (c *Client) DataTypeToBigQuerySchema(dataType arrow.DataType) bigquery.Schema {
	// handle complex types
	switch dataType.ID() {
	case arrow.STRUCT:
		switch v := dataType.(type) {
		case *arrow.StructType:
			fields := make([]*bigquery.FieldSchema, len(v.Fields()))
			for i, field := range v.Fields() {
				fields[i] = c.ColumnToBigQuerySchema(schema.Column{
					Name: field.Name,
					Type: field.Type,
				})
			}
			return fields
		}
	case arrow.LIST:
		switch v := dataType.(type) {
		case *arrow.ListType:
			return c.DataTypeToBigQuerySchema(v.Elem())
		case *arrow.LargeListType:
			return c.DataTypeToBigQuerySchema(v.Elem())
		}
	}

	// handle basic types
	switch {
	case arrow.TypeEqual(dataType, arrow.FixedWidthTypes.MonthInterval):
		return []*bigquery.FieldSchema{
			{
				Name: "months",
				Type: bigquery.IntegerFieldType,
			},
		}
	case arrow.TypeEqual(dataType, arrow.FixedWidthTypes.DayTimeInterval):
		return []*bigquery.FieldSchema{
			{
				Name: "days",
				Type: bigquery.IntegerFieldType,
			},
			{
				Name: "milliseconds",
				Type: bigquery.IntegerFieldType,
			},
		}
	case arrow.TypeEqual(dataType, arrow.FixedWidthTypes.MonthDayNanoInterval):
		return []*bigquery.FieldSchema{
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
	case typeOneOf(dataType, arrow.FixedWidthTypes.Timestamp_ns):
		return []*bigquery.FieldSchema{
			{
				Name: "timestamp",
				Type: bigquery.TimestampFieldType,
			},
			{
				Name: "nanoseconds",
				Type: bigquery.IntegerFieldType,
			},
		}
	case typeOneOf(dataType, arrow.FixedWidthTypes.Duration_ns):
		return []*bigquery.FieldSchema{
			{
				Name: "duration",
				Type: bigquery.IntervalFieldType,
			},
			{
				Name: "nanoseconds",
				Type: bigquery.IntegerFieldType,
			},
		}
	}
	return nil
}

func typeOneOf(left arrow.DataType, dt ...arrow.DataType) bool {
	for _, t := range dt {
		if arrow.TypeEqual(left, t) {
			return true
		}
	}
	return false
}
