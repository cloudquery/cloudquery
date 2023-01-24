package mixpanel

import (
	"bytes"
	"time"
)

type Time time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *Time) MarshalJSON() ([]byte, error) {
	if t == nil || time.Time(*t).IsZero() {
		return []byte("null"), nil
	}
	u := time.Time(*t).Format(timeFormat)
	return []byte(u), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		tt := Time(time.Time{})
		*t = tt
		return nil
	}

	tt, err := time.Parse(timeFormat, string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}
