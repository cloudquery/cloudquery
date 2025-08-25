package client

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

var (
	reTimestamp = regexp.MustCompile(`timestamp(?:_(?:ltz|tz|ntz))?\s*(?:\(([0-9])\))?`)
	reTime      = regexp.MustCompile(`time\s*(?:\(([0-9])\))?`)
	reNumeric   = regexp.MustCompile(`(?:numeric|number|decimal)\s*(?:\(([0-9]+)\s*(?:,\s*([0-9]+))?\))?`)
)

func SchemaTypeToSnowflake(t arrow.DataType) string {
	switch t.(type) {
	case *arrow.ListType, *arrow.FixedSizeListType:
		return "array"
	case *arrow.BooleanType:
		return "boolean"
	case *arrow.Int8Type, *arrow.Uint8Type, *arrow.Int16Type, *arrow.Uint16Type,
		*arrow.Int32Type, *arrow.Uint32Type, *arrow.Int64Type, *arrow.Uint64Type:
		return "number"
	case *arrow.Float32Type, *arrow.Float64Type:
		return "float"
	case *arrow.StringType, *arrow.LargeStringType:
		return "text"
	case *arrow.BinaryType, *arrow.LargeBinaryType:
		return "binary"
	case *arrow.TimestampType:
		return "timestamp_tz"
	case *types.JSONType, *arrow.StructType:
		return "variant"
	default:
		return "text"
	}
}

func SnowflakeToSchemaType(t string) arrow.DataType {
	t = strings.ToLower(strings.TrimSpace(t))
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(SnowflakeToSchemaType(t[:len(t)-2]))
	}

	// Try specialized parsers first
	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
		parseNumeric,
	}
	for _, parser := range parsers {
		if got, matched := parser(t); matched {
			return got
		}
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean

	case "tinyint", "smallint", "integer", "int", "bigint":
		return arrow.PrimitiveTypes.Int64

	case "float", "float4", "float8", "double", "double precision", "real":
		return arrow.PrimitiveTypes.Float64

	case "date":
		return arrow.FixedWidthTypes.Date32

	case "char", "character", "varchar", "string", "text",
		"geography", "geometry":
		return arrow.BinaryTypes.String

	case "binary", "varbinary":
		return arrow.BinaryTypes.Binary

	case "variant", "object", "array":
		return types.ExtensionTypes.JSON

	default:
		return arrow.BinaryTypes.String
	}
}

func parseTimestamp(t string) (arrow.DataType, bool) {
	// Handle Snowflake TIMESTAMP_* types with optional precision
	if t == "timestamp_ltz" || t == "timestamp_tz" || t == "timestamp_ntz" {
		return arrow.FixedWidthTypes.Timestamp_ns, true
	}

	matches := reTimestamp.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}

	precisionStr := matches[0][1]
	precision := 9 // default
	if precisionStr != "" {
		p, err := strconv.Atoi(precisionStr)
		if err == nil {
			precision = p
		}
	}

	switch {
	case precision == 0:
		return arrow.FixedWidthTypes.Timestamp_s, true
	case precision >= 1 && precision <= 3:
		return arrow.FixedWidthTypes.Timestamp_ms, true
	case precision >= 4 && precision <= 6:
		return arrow.FixedWidthTypes.Timestamp_us, true
	case precision >= 7 && precision <= 9:
		return arrow.FixedWidthTypes.Timestamp_ns, true
	default:
		return arrow.FixedWidthTypes.Timestamp_ns, true
	}
}

func parseTime(t string) (arrow.DataType, bool) {
	matches := reTime.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}

	precisionStr := matches[0][1]
	precision := 9 // default
	if precisionStr != "" {
		p, err := strconv.Atoi(precisionStr)
		if err == nil {
			precision = p
		}
	}

	switch {
	case precision == 0:
		return arrow.FixedWidthTypes.Time32s, true
	case precision >= 1 && precision <= 3:
		return arrow.FixedWidthTypes.Time32ms, true
	case precision >= 4 && precision <= 6:
		return arrow.FixedWidthTypes.Time64us, true
	case precision >= 7 && precision <= 9:
		return arrow.FixedWidthTypes.Time64ns, true
	default:
		return arrow.FixedWidthTypes.Time64ns, true
	}
}
func parseNumeric(t string) (arrow.DataType, bool) {
	matches := reNumeric.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}

	// No precision/scale specified - default to Int64
	if len(matches[0]) < 3 || matches[0][1] == "" {
		return arrow.PrimitiveTypes.Int64, true
	}

	precision, err := strconv.ParseInt(matches[0][1], 10, 32)
	if precision == 0 || err != nil {
		panic("precision cannot be 0")
	}
	scale, err := strconv.ParseInt(matches[0][2], 10, 32)
	if err != nil {
		panic("error parsing scale " + err.Error())
	}

	switch {
	case precision <= 18 && scale == 0:
		return arrow.PrimitiveTypes.Int64, true
	case precision == 20 && scale == 0:
		return arrow.PrimitiveTypes.Uint64, true
	case precision == 38 && scale == 0:
		return arrow.PrimitiveTypes.Int64, true
	case precision <= 38:
		return &arrow.Decimal128Type{Precision: int32(precision), Scale: int32(scale)}, true
	case precision <= 76:
		return &arrow.Decimal256Type{Precision: int32(precision), Scale: int32(scale)}, true
	default:
		return arrow.BinaryTypes.String, true
	}
}
