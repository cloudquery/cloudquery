package vercel

import (
	"strconv"
	"time"
)

type MilliTime time.Time

func (t *MilliTime) MarshalJSON() ([]byte, error) {
	if t == nil || time.Time(*t).IsZero() {
		return []byte("null"), nil
	}
	u := time.Time(*t).UnixMilli()
	return []byte(strconv.FormatInt(u, 10)), nil
}

func (t *MilliTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		tt := MilliTime(time.Time{})
		*t = tt
		return nil
	}

	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = MilliTime(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}
