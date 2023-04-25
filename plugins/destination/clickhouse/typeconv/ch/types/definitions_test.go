package types

import (
	"strings"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/stretchr/testify/require"
)

func TestFieldType(t *testing.T) {
	for _, tc := range []testCase{
		{_type: new(arrow.BooleanType), expected: "Bool"},
		{_type: new(arrow.Int8Type), expected: "Int8"},
		{_type: new(arrow.Int16Type), expected: "Int16"},
		{_type: new(arrow.Int32Type), expected: "Int32"},
		{_type: new(arrow.Int64Type), expected: "Int64"},
		{_type: new(arrow.Uint8Type), expected: "UInt8"},
		{_type: new(arrow.Uint16Type), expected: "UInt16"},
		{_type: new(arrow.Uint32Type), expected: "UInt32"},
		{_type: new(arrow.Uint64Type), expected: "UInt64"},
		{_type: new(arrow.Float16Type), expected: "Float32"},
		{_type: new(arrow.Float32Type), expected: "Float32"},
		{_type: new(arrow.Float64Type), expected: "Float64"},
		{_type: &arrow.FixedSizeBinaryType{ByteWidth: 125}, expected: "FixedString(125)"},
		{_type: new(arrow.Date32Type), expected: "Date32"},
		{_type: new(arrow.Date64Type), expected: "DateTime64(3)"},
		{_type: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)), expected: "String"},
	} {
		ensureDefinition(t, tc)
	}
}

type testCase struct {
	_type    arrow.DataType
	expected string
}

func ensureDefinition(t *testing.T, tc testCase) {
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

	t.Run(tc.expected, func(t *testing.T) {
		// non-nullable
		field := arrow.Field{
			Name: replacer.Replace(tc._type.String()),
			Type: tc._type,
		}
		fieldType, err := FieldType(field)
		require.NoError(t, err)
		require.Equal(t, tc.expected, fieldType)

		if field.Type.ID() == arrow.LIST {
			// arrays cannot be marked nullable in ClickHouse
			return
		}

		// nullable
		field.Nullable = true
		fieldType, err = FieldType(field)
		require.NoError(t, err)
		require.Equal(t, "Nullable("+tc.expected+")", fieldType)
	})
}
