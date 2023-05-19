package changes

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	schemav2 "github.com/cloudquery/plugin-sdk/v2/schema"
)

var cqToArrow = map[string]arrow.DataType{}

func init() {
	for t := schemav2.TypeInvalid + 1; t < schemav2.TypeTimeIntervalDeprecated; t++ {
		cqToArrow[strings.TrimPrefix(t.String(), "Type")] = schemav2.CQColumnToArrowField(&schemav2.Column{Type: t}).Type
	}
}

// dataTypesEqual checks if the old & current are the same OR if the current is an Arrow mapping of the old
func dataTypesEqual(old, current string) bool {
	if old == current {
		return true
	}

	dt, ok := cqToArrow[old]
	if !ok {
		return false
	}

	return dt.String() == current
}
