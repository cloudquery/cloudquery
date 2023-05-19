package changes

import (
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	schemav2 "github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

var cqToArrow = map[string]arrow.DataType{}

func init() {
	for t := schemav2.TypeInvalid + 1; t < schemav2.TypeTimeIntervalDeprecated; t++ {
		cqToArrow[strings.TrimPrefix(t.String(), "Type")] = schemav2.CQColumnToArrowField(&schemav2.Column{Type: t}).Type
	}

	// extensions are special as we need v3/types for extensions (diff in `String` func)
	cqToArrow[strings.TrimPrefix(schemav2.TypeUUID.String(), "Type")] = types.ExtensionTypes.UUID
	cqToArrow[strings.TrimPrefix(schemav2.TypeInet.String(), "Type")] = types.ExtensionTypes.Inet
	cqToArrow[strings.TrimPrefix(schemav2.TypeMacAddr.String(), "Type")] = types.ExtensionTypes.MAC
	cqToArrow[strings.TrimPrefix(schemav2.TypeJSON.String(), "Type")] = types.ExtensionTypes.JSON
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
