package transformers

import (
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/recordupdater"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/spec"
	"github.com/cloudquery/cloudquery/plugins/transformer/basic/client/tablematcher"
)

type TransformationFn = func(arrow.Record) (arrow.Record, error)

type Transformer struct {
	matcher *tablematcher.TableMatcher
	fn      TransformationFn
}

func NewFromSpec(sp spec.TransformationSpec) (*Transformer, error) {
	tr := &Transformer{matcher: tablematcher.New(sp.Tables)}

	switch sp.Kind {
	case spec.KindAddColumn:
		tr.fn = AddLiteralStringColumnAsLastColumn(sp.Name, sp.Value)
	case spec.KindRemoveColumns:
		tr.fn = RemoveColumns(sp.Columns)
	case spec.KindObfuscateColumns:
		tr.fn = ObfuscateColumns(sp.Columns)
	default:
		return nil, fmt.Errorf("unknown transformation kind: %s", sp.Kind)
	}

	return tr, nil
}

func (tr *Transformer) Transform(record arrow.Record) (arrow.Record, error) {
	// Passthrough if the record's table is not a match to any of the spec's
	// tablepatterns, but error if the record doesn't have table metadata.
	isMatch, err := tr.matcher.IsRecordsTableMatch(record)
	if err != nil {
		return nil, err
	}
	if !isMatch {
		return record, nil
	}
	// Apply the specific transformation kind
	return tr.fn(record)
}

func AddLiteralStringColumnAsLastColumn(name, value string) TransformationFn {
	return func(record arrow.Record) (arrow.Record, error) {
		return recordupdater.New(record).AddLiteralStringColumn(name, value, -1)
	}
}

func RemoveColumns(columnNames []string) TransformationFn {
	return func(record arrow.Record) (arrow.Record, error) {
		return recordupdater.New(record).RemoveColumns(columnNames)
	}
}

func ObfuscateColumns(columnNames []string) TransformationFn {
	return func(record arrow.Record) (arrow.Record, error) {
		return recordupdater.New(record).ObfuscateColumns(columnNames)
	}
}
