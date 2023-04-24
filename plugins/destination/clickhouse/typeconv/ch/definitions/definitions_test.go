package definitions

import (
	"strings"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/stretchr/testify/require"
)

func TestFieldType(t *testing.T) {
	type testCase struct {
		data     arrow.DataType
		expected string
	}

	for _, tc := range []testCase{
		{data: new(arrow.BooleanType), expected: "Bool"},
		{data: new(arrow.Int8Type), expected: "Int8"},
		{data: new(arrow.Int16Type), expected: "Int16"},
		{data: new(arrow.Int32Type), expected: "Int32"},
		{data: new(arrow.Int64Type), expected: "Int64"},
		{data: new(arrow.Uint8Type), expected: "UInt8"},
		{data: new(arrow.Uint16Type), expected: "UInt16"},
		{data: new(arrow.Uint32Type), expected: "UInt32"},
		{data: new(arrow.Uint64Type), expected: "UInt64"},
		{data: new(arrow.Float16Type), expected: "Float32"},
		{data: new(arrow.Float32Type), expected: "Float32"},
		{data: new(arrow.Float64Type), expected: "Float64"},
		{data: &arrow.FixedSizeBinaryType{ByteWidth: 125}, expected: "FixedString(125)"},
		{data: new(arrow.Date32Type), expected: "Date32"},
		{data: new(arrow.Date64Type), expected: "DateTime64(3)"},
		{data: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)), expected: "String"},
	} {
		ensureDefinition(t, tc.data, tc.expected)
	}
}

func ensureDefinition(t *testing.T, _type arrow.DataType, expected string) {
	replacer := strings.NewReplacer(
		"@", "_",
		"(", "_",
		"<", "_",
		"[", "_",
		")", "_",
		">", "_",
		"]", "_",
		" ", "_",
		":", "_",
		",", "_",
	)

	// non-nullable
	field := arrow.Field{
		Name: replacer.Replace(_type.String()),
		Type: _type,
	}
	fieldType := FieldType(field)
	require.Equal(t, expected, fieldType)

	if field.Type.ID() == arrow.LIST {
		// arrays cannot be marked nullable in ClickHouse
		return
	}

	// nullable
	field.Nullable = true
	fieldType = FieldType(field)
	require.Equal(t, "Nullable("+expected+")", fieldType)
}
