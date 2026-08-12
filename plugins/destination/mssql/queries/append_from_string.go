package queries

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
	default:
		return b.AppendValueFromString(s)
	}
}
