package values

import (
	"testing"
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/stretchr/testify/require"
)

func Test_timestamp(t *testing.T) {
	const (
		amount = 100
		unit   = arrow.Millisecond
	)
	for i := 0; i < amount; i++ {
		ts := time.Now().UTC().AddDate(0, i, i).Add(time.Duration(i * 123))
		t.Run(ts.String(), func(t *testing.T) {
			aTs, err := arrow.TimestampFromTime(ts, unit)
			require.NoError(t, err)
			ensureRecord(t, testCase{dataType: &arrow.TimestampType{Unit: arrow.Millisecond}, value: &ts, expected: aTs})
		})
	}
}
