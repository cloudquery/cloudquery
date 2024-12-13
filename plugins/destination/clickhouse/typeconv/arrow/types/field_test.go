package types

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/stretchr/testify/require"
)

func TestField(t *testing.T) {
	for _, tc := range []testCase{
		{columnType: "Bool", expected: new(arrow.BooleanType)},
		{columnType: "Int8", expected: new(arrow.Int8Type)},
		{columnType: "Int16", expected: new(arrow.Int16Type)},
		{columnType: "Int32", expected: new(arrow.Int32Type)},
		{columnType: "Int64", expected: new(arrow.Int64Type)},
		{columnType: "UInt8", expected: new(arrow.Uint8Type)},
		{columnType: "UInt16", expected: new(arrow.Uint16Type)},
		{columnType: "UInt32", expected: new(arrow.Uint32Type)},
		{columnType: "UInt64", expected: new(arrow.Uint64Type)},
		{columnType: "Float32", expected: new(arrow.Float32Type)},
		{columnType: "Float64", expected: new(arrow.Float64Type)},
		{columnType: "FixedString(125)", expected: &arrow.FixedSizeBinaryType{ByteWidth: 125}},
		{columnType: "Date", expected: new(arrow.Date32Type)},
		{columnType: "Date32", expected: new(arrow.Date32Type)},
		{columnType: "UUID", expected: new(types.UUIDType)},
	} {
		ensureField(t, tc)
	}
}

type testCase struct {
	columnType string
	expected   arrow.DataType
}

func ensureField(t *testing.T, tc testCase) {
	t.Helper()
	t.Run(tc.columnType, func(t *testing.T) {
		t.Helper()
		// simple
		field, err := Field("field", tc.columnType)
		require.NoError(t, err)
		require.NotNil(t, field)
		require.Truef(t, arrow.TypeEqual(tc.expected, field.Type), "expected type:\n%s\nactual:\n%s", tc.expected.String(), field.Type.String())
		require.False(t, field.Nullable)

		// nullable
		field, err = Field("field", "Nullable("+tc.columnType+")")
		require.NoError(t, err)
		require.True(t, field.Nullable)
		require.True(t, arrow.TypeEqual(tc.expected, field.Type))
	})
}
