package values

import (
	"testing"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/decimal128"
	"github.com/apache/arrow/go/v14/arrow/decimal256"
	"github.com/apache/arrow/go/v14/arrow/float16"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	dataType arrow.DataType
	value    any
	expected any
}

func ensureRecord(t *testing.T, tc testCase) {
	t.Helper()
	t.Run(tc.dataType.String(), func(t *testing.T) {
		t.Helper()
		table := &schema.Table{Columns: schema.ColumnList{{Name: "field", Type: tc.dataType}}}
		record, err := Record(table.ToArrowSchema(), []any{tc.value})
		require.NoError(t, err)
		require.Equal(t, int64(1), record.NumRows())
		require.Equal(t, int64(1), record.NumCols())

		column := record.Column(0)
		require.True(t, column.IsValid(0))

		require.EqualValues(t, tc.expected, getValue0(column))
	})
}

func getValue0(arr arrow.Array) any {
	type valuer[A any] interface {
		Value(int) A
	}
	switch arr := arr.(type) {
	case valuer[bool]:
		return arr.Value(0)

	case valuer[uint8]:
		return arr.Value(0)
	case valuer[uint16]:
		return arr.Value(0)
	case valuer[uint32]:
		return arr.Value(0)
	case valuer[uint64]:
		return arr.Value(0)

	case valuer[int8]:
		return arr.Value(0)
	case valuer[int16]:
		return arr.Value(0)
	case valuer[int32]:
		return arr.Value(0)
	case valuer[int64]:
		return arr.Value(0)

	case valuer[float16.Num]:
		return arr.Value(0)
	case valuer[float32]:
		return arr.Value(0)
	case valuer[float64]:
		return arr.Value(0)

	case valuer[[]byte]:
		return arr.Value(0)
	case valuer[string]:
		return arr.Value(0)

	case valuer[uuid.UUID]:
		return arr.Value(0)

	case valuer[decimal128.Num]:
		return arr.Value(0)
	case valuer[decimal256.Num]:
		return arr.Value(0)

	case valuer[arrow.Time32]:
		return arr.Value(0)
	case valuer[arrow.Time64]:
		return arr.Value(0)
	case valuer[arrow.Date32]:
		return arr.Value(0)
	case valuer[arrow.Date64]:
		return arr.Value(0)
	case valuer[arrow.Timestamp]:
		return arr.Value(0)

	default:
		return arr.GetOneForMarshal(0)
	}
}
