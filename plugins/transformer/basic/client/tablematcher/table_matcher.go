package tablematcher

import (
	"errors"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/plugin-sdk/v4/glob"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

// TableMatcher takes a slice of glob patterns and answers whether an `arrow.Record`
// belongs to a table matching any of the patterns (via Schema().Metadata()).
//
// It is used to allowlist records for transformations, based on spec.
//
// Unlike sources, which know which tables exist, the transformer discovers
// existing tables by observing records. Thus, a cache is used to keep constant lookup time.
type TableMatcher struct {
	patterns     []string
	matcherCache map[string]bool
}

func New(patterns []string) *TableMatcher {
	return &TableMatcher{
		patterns:     patterns,
		matcherCache: make(map[string]bool),
	}
}

func (t *TableMatcher) isTableMatch(tableName string) bool {
	if result, ok := t.matcherCache[tableName]; ok {
		return result
	}
	isMatch := false
	for _, p := range t.patterns {
		if glob.Glob(p, tableName) {
			isMatch = true
			break
		}
	}
	t.matcherCache[tableName] = isMatch
	return isMatch
}

func (t *TableMatcher) IsSchemasTableMatch(sc *arrow.Schema) (bool, error) {
	table, ok := sc.Metadata().GetValue(schema.MetadataTableName)
	if !ok {
		return false, errors.New("table name not found in record's metadata")
	}
	return t.isTableMatch(table), nil
}
