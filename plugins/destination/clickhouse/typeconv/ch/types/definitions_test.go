package types

import (
	"strings"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/stretchr/testify/require"
)

func TestFieldType(t *testing.T) {
	for _, tc := range []testCase{
		{dataType: new(arrow.BooleanType), expected: "Bool"},
		{dataType: new(arrow.Int8Type), expected: "Int8"},
		{dataType: new(arrow.Int16Type), expected: "Int16"},
		{dataType: new(arrow.Int32Type), expected: "Int32"},
		{dataType: new(arrow.Int64Type), expected: "Int64"},
		{dataType: new(arrow.Uint8Type), expected: "UInt8"},
		{dataType: new(arrow.Uint16Type), expected: "UInt16"},
		{dataType: new(arrow.Uint32Type), expected: "UInt32"},
		{dataType: new(arrow.Uint64Type), expected: "UInt64"},
		{dataType: new(arrow.Float16Type), expected: "Float32"},
		{dataType: new(arrow.Float32Type), expected: "Float32"},
		{dataType: new(arrow.Float64Type), expected: "Float64"},
		{dataType: &arrow.FixedSizeBinaryType{ByteWidth: 125}, expected: "FixedString(125)"},
		{dataType: new(arrow.Date32Type), expected: "Date32"},
		{dataType: new(arrow.Date64Type), expected: "DateTime"},
		{dataType: &arrow.Time32Type{Unit: arrow.Second}, expected: "DateTime64(0)"},
		{dataType: &arrow.Time32Type{Unit: arrow.Millisecond}, expected: "DateTime64(3)"},
		{dataType: &arrow.Time64Type{Unit: arrow.Microsecond}, expected: "DateTime64(6)"},
		{dataType: &arrow.Time64Type{Unit: arrow.Nanosecond}, expected: "DateTime64(9)"},
	} {
		ensureDefinition(t, tc)
	}
}

type testCase struct {
	dataType arrow.DataType
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

	t.Run(tc.dataType.String(), func(t *testing.T) {
		// non-nullable
		col := schema.Column{
			Name:    replacer.Replace(tc.dataType.String()),
			Type:    tc.dataType,
			NotNull: true,
		}
		typeDef, err := FieldType(col.ToArrowField())
		require.NoError(t, err)
		require.Equal(t, tc.expected, typeDef)

		if !CanBeNullable(col.Type) {
			return
		}

		// nullable
		col.NotNull = false
		typeDef, err = FieldType(col.ToArrowField())
		require.NoError(t, err)
		require.Equal(t, "Nullable("+tc.expected+")", typeDef)
	})
}
