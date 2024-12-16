package values

import (
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
)

func Test_date32(t *testing.T) {
	for _, date := range []arrow.Date32{0, 1, 2, 3, 100, 1000, 10000, 100000} {
		val := date.ToTime()
		ensureRecord(t, testCase{dataType: new(arrow.Date32Type), value: &val, expected: date})
	}
}

func Test_date64(t *testing.T) {
	const msInHour = int64(time.Hour / time.Millisecond)
	for _, hours := range []int64{0, 1, 2, 3, 100, 1000, 10000, 100000} {
		date := arrow.Date64(hours * msInHour)
		val := date.ToTime()
		date = arrow.Date64FromTime(val)
		ensureRecord(t, testCase{dataType: new(arrow.Date64Type), value: &val, expected: date})
	}
}
