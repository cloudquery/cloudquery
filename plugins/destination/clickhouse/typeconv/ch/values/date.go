package values

import (
	"time"
)

type date interface {
	ToTime() time.Time
}

func dateValue[A date, ARR primitive[A]](arr ARR) []*time.Time {
	res := make([]*time.Time, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsValid(i) {
			val := arr.Value(i).ToTime()
			res[i] = &val
		}
	}
	return res
}
