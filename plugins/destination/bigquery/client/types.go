package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

// TimestampNanoseconds is a struct to hold a timestamp with nanosecond precision,
// because BigQuery does not support nanosecond precision timestamps.
type TimestampNanoseconds struct {
	Timestamp   string `json:"timestamp" bigquery:"timestamp"`
	Nanoseconds int    `json:"nanoseconds" bigquery:"nanoseconds"`
}

func (c *Client) ColumnToBigQuerySchema(col schema.Column) *bigquery.FieldSchema {
	sc := bigquery.FieldSchema{
		Name:        col.Name,
		Description: col.Description,
		Type:        c.DataTypeToBigQueryType(col.Type),
		Schema:      c.DataTypeToBigQuerySchema(col.Type),
	}

	// TODO: handle repeated; we currently don't handle the case where we get a list of lists,
	//       but that's not a valid case right now and isn't being explicitly tested for.
	if isListType(col.Type) {
		sc.Repeated = true
	}
	return &sc
}

func isListType(t arrow.DataType) bool {
	return t.ID() == arrow.LIST || t.ID() == arrow.LARGE_LIST || t.ID() == arrow.FIXED_SIZE_LIST
}

func (c *Client) DataTypeToBigQueryType(dataType arrow.DataType) bigquery.FieldType {
	switch {
	// handle known extensions that require special handling
	case typeOneOf(dataType,
		types.ExtensionTypes.JSON):
		return bigquery.JSONFieldType
	case typeOneOf(dataType,
		types.ExtensionTypes.Inet,
		types.ExtensionTypes.MAC,
		types.ExtensionTypes.UUID):
		return bigquery.StringFieldType

	// handle complex types
	case dataType.ID() == arrow.MAP:
		return bigquery.JSONFieldType
	case dataType.ID() == arrow.STRUCT:
		return bigquery.RecordFieldType
	case isListType(dataType):
		switch v := dataType.(type) {
		case *arrow.ListType:
			return c.DataTypeToBigQueryType(v.Elem())
		case *arrow.LargeListType:
			return c.DataTypeToBigQueryType(v.Elem())
		case *arrow.FixedSizeListType:
			return c.DataTypeToBigQueryType(v.Elem())
		}
		fallthrough

	// handle basic types
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
		arrow.PrimitiveTypes.Uint32):
		return bigquery.IntegerFieldType
	case typeOneOf(dataType,
		arrow.PrimitiveTypes.Uint64):
		return bigquery.NumericFieldType
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
		arrow.FixedWidthTypes.Duration_us,
		arrow.FixedWidthTypes.Duration_ns):
		// BigQuery does not support intervals with precisions higher than seconds,
		// and in the case of seconds the max value is not large enough to contain the
		// max Arrow duration, so we store durations as plain integers. Users will need
		// to cast or transform to interval, if necessary.
		return bigquery.IntegerFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.MonthInterval):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.DayTimeInterval):
		return bigquery.RecordFieldType
	case typeOneOf(dataType,
		arrow.FixedWidthTypes.MonthDayNanoInterval):
		return bigquery.RecordFieldType
		// We don't use `typeOneOf` as `arrow.TypeEqual` checks for equality of precision and scale.
	case arrow.IsDecimal(dataType.ID()):
		// BigQuery NumericFieldType has a scale limit of 9, so we use BigNumeric for both decimal128 and decimal256.
		return bigquery.BigNumericFieldType
	default:
		panic("unsupported data type: " + dataType.String())
	}
}

func (c *Client) DataTypeToBigQuerySchema(dataType arrow.DataType) bigquery.Schema {
	switch {
	case dataType.ID() == arrow.STRUCT:
		v := dataType.(*arrow.StructType)
		fields := make([]*bigquery.FieldSchema, len(v.Fields()))
		for i, field := range v.Fields() {
			fields[i] = c.ColumnToBigQuerySchema(schema.Column{
				Name: field.Name,
				Type: field.Type,
			})
		}
		return fields
	case arrow.IsListLike(dataType.ID()):
		switch v := dataType.(type) {
		case *arrow.ListType:
			return c.DataTypeToBigQuerySchema(v.Elem())
		case *arrow.LargeListType:
			return c.DataTypeToBigQuerySchema(v.Elem())
		default:
			panic("unsupported list type: " + dataType.String())
		}
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
