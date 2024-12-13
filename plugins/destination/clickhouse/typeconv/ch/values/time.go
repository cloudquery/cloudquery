package values

import (
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
)

func timestampValue(arr *array.Timestamp) ([]*time.Time, error) {
	conv, err := arr.DataType().(*arrow.TimestampType).GetToTimeFunc()
	if err != nil {
		return nil, err
	}

	res := make([]*time.Time, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := conv(arr.Value(i))
			res[i] = &val
		}
	}

	return res, nil
}

type timeWithUnit interface {
	ToTime(unit arrow.TimeUnit) time.Time
}

func timeValue[A timeWithUnit, ARR primitive[A]](arr ARR, unit arrow.TimeUnit) []*time.Time {
	res := make([]*time.Time, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := arr.Value(i).ToTime(unit)
			res[i] = &val
		}
	}
	return res
}
