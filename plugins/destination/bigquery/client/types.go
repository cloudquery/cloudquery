package client

import (
	"cloud.google.com/go/bigquery"
	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

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

func isListType(dt arrow.DataType) bool {
	switch dt.(type) {
	case *arrow.MapType:
		return false
	case arrow.ListLikeType:
		return true
	default:
		return false
	}
}

func (c *Client) DataTypeToBigQueryType(dt arrow.DataType) bigquery.FieldType {
	switch dt := dt.(type) {
	// handle known extensions that require special handling
	case *types.JSONType:
		return bigquery.JSONFieldType
	case *types.InetType, *types.MACType, *types.UUIDType:
		return bigquery.StringFieldType

	// handle complex types
	case *arrow.MapType:
		return bigquery.JSONFieldType
	case *arrow.StructType:
		return bigquery.RecordFieldType
	case arrow.ListLikeType:
		return c.DataTypeToBigQueryType(dt.Elem())

	// handle basic types
	case *arrow.BooleanType:
		return bigquery.BooleanFieldType
	case *arrow.Int8Type,
		*arrow.Int16Type,
		*arrow.Int32Type,
		*arrow.Int64Type,
		*arrow.Uint8Type,
		*arrow.Uint16Type,
		*arrow.Uint32Type:
		return bigquery.IntegerFieldType
	case *arrow.Uint64Type:
		return bigquery.NumericFieldType
	case *arrow.Float16Type, *arrow.Float32Type, *arrow.Float64Type:
		return bigquery.FloatFieldType
	case arrow.BinaryDataType:
		if dt.IsUtf8() {
			return bigquery.StringFieldType
		}
		return bigquery.BytesFieldType
	case *arrow.Date32Type, *arrow.Date64Type:
		return bigquery.DateFieldType
	case *arrow.TimestampType:
		return bigquery.TimestampFieldType
	case *arrow.Time32Type, *arrow.Time64Type:
		return bigquery.TimeFieldType
	case *arrow.DurationType:
		// BigQuery does not support intervals with precisions higher than seconds,
		// and in the case of seconds the max value is not large enough to contain the
		// max Arrow duration, so we store durations as plain integers. Users will need
		// to cast or transform to interval, if necessary.
		return bigquery.IntegerFieldType
	case *arrow.MonthIntervalType,
		*arrow.DayTimeIntervalType,
		*arrow.MonthDayNanoIntervalType:
		return bigquery.RecordFieldType
		// We don't use `typeOneOf` as `arrow.TypeEqual` checks for equality of precision and scale.
	case arrow.DecimalType:
		// BigQuery NumericFieldType has a scale limit of 9, so we use BigNumeric for both decimal128 and decimal256.
		return bigquery.BigNumericFieldType
	default:
		panic("unsupported data type: " + dt.String())
	}
}

func (c *Client) DataTypeToBigQuerySchema(dt arrow.DataType) bigquery.Schema {
	switch dt := dt.(type) {
	case *arrow.StructType:
		fields := make([]*bigquery.FieldSchema, len(dt.Fields()))
		for i, field := range dt.Fields() {
			fields[i] = c.ColumnToBigQuerySchema(schema.Column{
				Name: field.Name,
				Type: field.Type,
			})
		}
		return fields
	case *arrow.MapType:
		return nil
	case arrow.ListLikeType:
		// TODO: handle maps
		return c.DataTypeToBigQuerySchema(dt.Elem())
	case *arrow.MonthIntervalType:
		return []*bigquery.FieldSchema{
			{
				Name: "months",
				Type: bigquery.IntegerFieldType,
			},
		}
	case *arrow.DayTimeIntervalType:
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
	case *arrow.MonthDayNanoIntervalType:
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
	}
	return nil
}
