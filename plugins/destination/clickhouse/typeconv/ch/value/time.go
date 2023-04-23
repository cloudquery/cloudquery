package value

import (
	"time"

	"github.com/apache/arrow/go/v12/arrow"
	"github.com/apache/arrow/go/v12/arrow/array"
)

type date interface {
	ToTime() time.Time
}

func dateValue[A date](arr primitive[A]) []*time.Time {
	res := make([]*time.Time, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := arr.Value(i).ToTime()
			res[i] = &val
		}
	}
	return res
}

func timestampValue(arr *array.Timestamp) ([]*time.Time, error) {
	conv, err := arr.DataType().(*arrow.TimestampType).GetToTimeFunc()
	if err != nil {
		return nil, err
	}

	res := make([]*time.Time, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) && !arr.IsNull(i) {
			val := conv(arr.Value(i))
			res[i] = &val
		}
	}

	return res, nil
}
