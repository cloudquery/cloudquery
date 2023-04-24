package definitions

import (
	"strings"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/stretchr/testify/require"
)

func TestFieldType(t *testing.T) {
	type testCase struct {
		data arrow.DataType
		exp  string
	}

	for _, tc := range []testCase{
		{data: new(arrow.BooleanType), exp: "Bool"},
		{data: new(arrow.Int8Type), exp: "Int8"},
		{data: new(arrow.Int16Type), exp: "Int16"},
		{data: new(arrow.Int32Type), exp: "Int32"},
		{data: new(arrow.Int64Type), exp: "Int64"},
		{data: new(arrow.Uint8Type), exp: "UInt8"},
		{data: new(arrow.Uint16Type), exp: "UInt16"},
		{data: new(arrow.Uint32Type), exp: "UInt32"},
		{data: new(arrow.Uint64Type), exp: "UInt64"},
		{data: new(arrow.Float16Type), exp: "Float32"},
		{data: new(arrow.Float32Type), exp: "Float32"},
		{data: new(arrow.Float64Type), exp: "Float64"},
		{data: &arrow.FixedSizeBinaryType{ByteWidth: 125}, exp: "FixedString(125)"},
		{data: new(arrow.Date32Type), exp: "Date32"},
		{data: new(arrow.Date64Type), exp: "DateTime64(3)"},
		{data: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)), exp: "String"},
	} {
		ensureDefinition(t, tc.data, tc.exp)
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
