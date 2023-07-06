package pgarrow

import (
	"regexp"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

var (
	reTimestamp = regexp.MustCompile(`timestamp\s*(?:\(([0-6])\))?(?: with(?:out)? time zone)?`)
	reTime      = regexp.MustCompile(`time\s*(?:\(([0-6])\))?(?: with(?:out)? time zone)?`)
)

func Pg10ToArrow(t string) arrow.DataType {
	t = normalize(t)
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(Pg10ToArrow(t[:len(t)-2]))
	}

	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
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
	case "smallserial":
		return arrow.PrimitiveTypes.Int16
	case "serial":
		return arrow.PrimitiveTypes.Int32
	case "bigserial", "serial8":
		return arrow.PrimitiveTypes.Int64
	case "smallint", "int2":
		return arrow.PrimitiveTypes.Int16
	case "integer", "int", "int4":
		return arrow.PrimitiveTypes.Int32
	case "bigint", "int8":
		return arrow.PrimitiveTypes.Int64
	case "double precision", "float8":
		return arrow.PrimitiveTypes.Float64
	case "real", "float4":
		return arrow.PrimitiveTypes.Float32
	case "uuid":
		return cqtypes.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "json", "jsonb":
		return cqtypes.ExtensionTypes.JSON
	case "cidr":
		return cqtypes.ExtensionTypes.Inet
	case "macaddr", "macaddr8":
		return cqtypes.ExtensionTypes.MAC
	case "inet":
		return cqtypes.ExtensionTypes.Inet
	case "date":
		return arrow.FixedWidthTypes.Date32
	default:
		return arrow.BinaryTypes.String
	}
}

func Pg10ToCockroach(t string) arrow.DataType {
	t = normalize(t)
	if strings.HasSuffix(t, "[]") {
		return arrow.ListOf(Pg10ToCockroach(t[:len(t)-2]))
	}

	parsers := []func(string) (arrow.DataType, bool){
		parseTimestamp,
		parseTime,
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
	case "serial2", "smallserial":
		return arrow.PrimitiveTypes.Int16
	case "serial4":
		return arrow.PrimitiveTypes.Int32
	case "serial8", "bigserial", "serial":
		return arrow.PrimitiveTypes.Int64
	case "smallint", "int2":
		return arrow.PrimitiveTypes.Int16
	case "int4":
		return arrow.PrimitiveTypes.Int32
	case "int", "bigint", "int8", "int64", "integer":
		// Cockroach has different aliases for ints
		return arrow.PrimitiveTypes.Int64
	case "double precision", "float8":
		return arrow.PrimitiveTypes.Float64
	case "real", "float4":
		return arrow.PrimitiveTypes.Float32
	case "uuid":
		return cqtypes.ExtensionTypes.UUID
	case "bytea":
		return arrow.BinaryTypes.Binary
	case "json", "jsonb":
		return cqtypes.ExtensionTypes.JSON
	case "cidr":
		return cqtypes.ExtensionTypes.Inet
	case "macaddr", "macaddr8":
		// Cockroach lacks MAC type
		return arrow.BinaryTypes.String
	case "inet":
		return cqtypes.ExtensionTypes.Inet
	case "date":
		return arrow.FixedWidthTypes.Date32
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
	case "1":
		return arrow.FixedWidthTypes.Timestamp_ms, true
	case "2":
		return arrow.FixedWidthTypes.Timestamp_ms, true
	case "3":
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
	case "1":
		return arrow.FixedWidthTypes.Time32ms, true
	case "2":
		return arrow.FixedWidthTypes.Time32ms, true
	case "3":
		return arrow.FixedWidthTypes.Time32ms, true
	default:
		return arrow.FixedWidthTypes.Time64us, true
	}
}
