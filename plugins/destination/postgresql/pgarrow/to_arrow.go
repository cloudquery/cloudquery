package pgarrow

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

var (
	reTimestamp = regexp.MustCompile(`timestamp\s*(?:\(([0-6])\))?(?: with(?:out)? time zone)?`)
	reTime      = regexp.MustCompile(`time\s*(?:\(([0-6])\))?(?: with(?:out)? time zone)?`)
	reNumeric   = regexp.MustCompile(`numeric\s*(?:\(([0-9]+)\s*,\s*([0-9]+)\))?`)
)

func Pg10ToArrow(t string) arrow.DataType {
	t = normalize(t)
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(Pg10ToArrow(t[:len(t)-2]))
	}

	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
		parseNumeric,
	}
	for _, parser := range parsers {
		got, matched := parser(t)
		if matched {
			return got
		}
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "smallserial", "smallint", "int2":
		return arrow.PrimitiveTypes.Int16
	case "serial", "integer", "int", "int4":
		return arrow.PrimitiveTypes.Int32
	case "bigserial", "serial8", "bigint", "int8":
		return arrow.PrimitiveTypes.Int64
	case "real", "float4":
		return arrow.PrimitiveTypes.Float32
	case "double precision", "float8":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return cqtypes.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "date":
		return arrow.FixedWidthTypes.Date32
	case "json", "jsonb":
		return cqtypes.ExtensionTypes.JSON
	case "cidr", "inet":
		return cqtypes.ExtensionTypes.Inet
	case "macaddr", "macaddr8":
		return cqtypes.ExtensionTypes.MAC
	default:
		return arrow.BinaryTypes.String
	}
}

func CockroachToArrow(t string) arrow.DataType {
	t = normalize(t)
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(CockroachToArrow(t[:len(t)-2]))
	}

	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
		parseNumeric,
	}
	for _, parser := range parsers {
		got, matched := parser(t)
		if matched {
			return got
		}
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "serial2", "smallserial", "smallint", "int2":
		return arrow.PrimitiveTypes.Int16
	case "serial4", "int4":
		return arrow.PrimitiveTypes.Int32
	// Cockroach has different aliases for ints
	case "serial8", "bigserial", "serial", "int", "bigint", "int8", "int64", "integer":
		return arrow.PrimitiveTypes.Int64
	case "real", "float4":
		return arrow.PrimitiveTypes.Float32
	case "double precision", "float8":
		return arrow.PrimitiveTypes.Float64
	case "uuid":
		return cqtypes.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "date":
		return arrow.FixedWidthTypes.Date32
	case "json", "jsonb":
		return cqtypes.ExtensionTypes.JSON
	case "cidr", "inet":
		return cqtypes.ExtensionTypes.Inet
	// Cockroach lacks MAC type
	// case "macaddr", "macaddr8":
	default:
		return arrow.BinaryTypes.String
	}
}

func CrateDBToArrow(t string) arrow.DataType {
	t = normalize(t)
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(Pg10ToArrow(t[:len(t)-2]))
	}

	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
		parseNumeric,
	}
	for _, parser := range parsers {
		got, matched := parser(t)
		if matched {
			return got
		}
	}

	switch t {
	case "boolean":
		return arrow.FixedWidthTypes.Boolean
	case "smallserial", "smallint", "int2":
		return arrow.PrimitiveTypes.Int16
	case "serial", "integer", "int", "int4":
		return arrow.PrimitiveTypes.Int32
	case "bigserial", "serial8", "bigint", "int8":
		return arrow.PrimitiveTypes.Int64
	case "real", "float4":
		return arrow.PrimitiveTypes.Float32
	case "double precision", "float8":
		return arrow.PrimitiveTypes.Float64
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "date":
		return arrow.FixedWidthTypes.Date32
	case "json", "jsonb", "object":
		return cqtypes.ExtensionTypes.JSON
	case "cidr", "inet", "ip":
		return cqtypes.ExtensionTypes.Inet
	// CrateDB does not support these types
	// case "macaddr", "macaddr8", "uuid":
	default:
		return arrow.BinaryTypes.String
	}
}

func normalize(t string) string {
	return strings.ToLower(strings.TrimSpace(t))
}

func parseTimestamp(t string) (arrow.DataType, bool) {
	timestamptzPrefix := "timestamptz using"
	t = strings.TrimPrefix(t, timestamptzPrefix)
	if t == "timestamptz" {
		return arrow.FixedWidthTypes.Timestamp_us, true
	}

	matches := reTimestamp.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}
	switch matches[0][1] {
	case "0":
		return arrow.FixedWidthTypes.Timestamp_s, true
	case "1", "2", "3":
		return arrow.FixedWidthTypes.Timestamp_ms, true
	default:
		return arrow.FixedWidthTypes.Timestamp_us, true
	}
}

func parseTime(t string) (arrow.DataType, bool) {
	matches := reTime.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}
	switch matches[0][1] {
	case "0":
		return arrow.FixedWidthTypes.Time32s, true
	case "1", "2", "3":
		return arrow.FixedWidthTypes.Time32ms, true
	default:
		return arrow.FixedWidthTypes.Time64us, true
	}
}

func parseNumeric(t string) (arrow.DataType, bool) {
	matches := reNumeric.FindAllStringSubmatch(t, -1)
	if len(matches) == 0 {
		return nil, false
	}

	if len(matches[0]) < 3 || matches[0][1] == "" {
		// no precision/scale specified
		return &arrow.Decimal128Type{Precision: 38, Scale: 0}, true
	}

	precision, err := strconv.ParseInt(matches[0][1], 10, 32)
	if precision == 0 || err != nil {
		panic("precision cannot be 0")
	}
	scale, err := strconv.ParseInt(matches[0][2], 10, 32)
	if err != nil {
		panic("error parsing scale " + err.Error())
	}

	if precision == 20 && scale == 0 {
		// special case
		return arrow.PrimitiveTypes.Uint64, true
	}

	if precision <= 38 {
		return &arrow.Decimal128Type{Precision: int32(precision), Scale: int32(scale)}, true
	}
	if precision <= 76 {
		return &arrow.Decimal256Type{Precision: int32(precision), Scale: int32(scale)}, true
	}

	return arrow.BinaryTypes.String, true
}
