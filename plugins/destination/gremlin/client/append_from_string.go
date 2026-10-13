package client

import (
	"strings"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/goccy/go-json"
)

// appendFromString appends a value from its string representation. Nested types are
// JSON-encoded, so they are decoded with UseNumber to keep int64/uint64 values that
// exceed float64 precision intact.
func appendFromString(b array.Builder, s string) error {
	switch b.Type().ID() {
	case arrow.LIST, arrow.LARGE_LIST, arrow.LIST_VIEW, arrow.LARGE_LIST_VIEW, arrow.FIXED_SIZE_LIST, arrow.MAP, arrow.STRUCT:
		dec := json.NewDecoder(strings.NewReader(s))
		dec.UseNumber()
		return b.UnmarshalOne(dec)
	case arrow.TIMESTAMP:
		return appendTimestampFromString(b.(*array.TimestampBuilder), s)
	default:
		return b.AppendValueFromString(s)
	}
}

func appendTimestampFromString(b *array.TimestampBuilder, s string) error {
	dt := b.Type().(*arrow.TimestampType)
	loc, err := dt.GetZone()
	if err != nil {
		return err
	}
	ts, _, err := arrow.TimestampFromStringInLocation(s, dt.Unit, loc)
	if err != nil {
		return err
	}
	b.Append(ts)
	return nil
}
