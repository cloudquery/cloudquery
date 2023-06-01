package pgarrow

import (
	"regexp"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	cqtypes "github.com/cloudquery/plugin-sdk/v3/types"
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
	case "smallint":
		return arrow.PrimitiveTypes.Int16
	case "smallserial":
		return arrow.PrimitiveTypes.Int16
	case "serial":
		return arrow.PrimitiveTypes.Int32
	case "integer", "int", "int4":
		return arrow.PrimitiveTypes.Int32
	case "bigint", "int8":
		return arrow.PrimitiveTypes.Int64
	case "bigserial", "serial8":
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
	return Pg10ToArrow(t)
}

func normalize(t string) string {
	return strings.ToLower(strings.TrimSpace(t))
}

func parseTimestamp(t string) (arrow.DataType, bool) {
	timestamptzPrefix := "timestamptz using"
	if strings.HasPrefix(t, timestamptzPrefix) {
		t = strings.TrimPrefix(t, timestamptzPrefix)
	}
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
