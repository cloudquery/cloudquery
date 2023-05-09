package values

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func Test_list(t *testing.T) {
	const amount = 100
	ensureRecord(t, genListTestCase(t, arrow.ListOf(new(arrow.StringType)), uuid.NewString, amount))
	ensureRecord(t, genListTestCase(t, arrow.ListOf(new(arrow.Float64Type)), rand.Float64, amount))
	ensureRecord(t, genListTestCase(t, arrow.ListOf(types.NewUUIDType()), uuid.New, amount))
}

func genListTestCase[A any](t *testing.T, _type arrow.DataType, _new func() A, n int) testCase {
	values := make([]*A, n)
	expected := make([]A, n)
	for i := range expected {
		val := _new()
		values[i], expected[i] = &val, val
	}

	return testCase{
		_type:    _type,
		value:    values,
		expected: marshalList(t, expected),
	}
}

func marshalList(t *testing.T, value any) json.RawMessage {
	data, err := json.Marshal(value)
	require.NoError(t, err)
	return data
}
