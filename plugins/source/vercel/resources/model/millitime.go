package model

import (
	"strconv"
	"time"
)

type MilliTime time.Time

func (t *MilliTime) MarshalJSON() ([]byte, error) {
	if t == nil || time.Time(*t).IsZero() {
		return []byte("null"), nil
	}
	return time.Time(*t).MarshalJSON()
}

func (t *MilliTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = MilliTime(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}
