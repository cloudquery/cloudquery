package values

import (
	"encoding/json"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/stretchr/testify/require"
)

func Test_map(t *testing.T) {
	for _, tc := range []testCase{
		{
			dataType: arrow.MapOf(new(arrow.StringType), new(arrow.BooleanType)),
			value:    map[string]any{"f1": true},
			expected: marshalMap(t, map[string]any{"f1": true}),
		},
		{
			dataType: arrow.MapOf(new(arrow.Float64Type), new(arrow.BooleanType)),
			value:    string(marshalMap(t, map[float64]any{5.3: true})),
			expected: marshalMap(t, map[float64]any{5.3: true}),
		},
		{
			dataType: arrow.MapOf(
				new(arrow.StringType),
				arrow.StructOf(
					arrow.Field{Name: "f", Type: new(arrow.BooleanType)},
					arrow.Field{Name: "f_nullable", Type: new(arrow.BooleanType), Nullable: true},
				),
			),
			value: map[string]any{"entry_key": map[string]any{
				"f":          true,
				"f_nullable": nil,
			}},
			expected: marshalMap(t, map[string]any{"entry_key": map[string]any{
				"f":          true,
				"f_nullable": nil,
			}}),
		},
	} {
		ensureRecord(t, tc)
	}
}

func marshalMap[K comparable, V any](t *testing.T, a map[K]V) json.RawMessage {
	flat := make([]map[string]any, len(a))
	i := 0
	for k, v := range a {
		flat[i] = map[string]any{"key": k, "value": v}
		i++
	}

	data, err := json.Marshal(flat)
	require.NoError(t, err)
	return data
}
