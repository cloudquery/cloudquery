package typeconv

import (
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/stretchr/testify/require"
)

func TestCanonizedFieldPreservesStructFieldNamesWithEscapes(t *testing.T) {
	field := arrow.Field{
		Name: "nested",
		Type: arrow.StructOf(
			arrow.Field{Name: `key\with\slash`, Type: new(arrow.BooleanType)},
			arrow.Field{Name: "key`tick", Type: new(arrow.BooleanType)},
		),
	}
	canonized, err := CanonizedField(field)
	require.NoError(t, err)
	require.Equal(t, field.Name, canonized.Name)
	require.True(t, arrow.TypeEqual(field.Type, canonized.Type))
}
