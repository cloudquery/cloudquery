package transformers

import (
	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/recordupdater"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/spec"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/tablematcher"
)

type TransformationFn = func(arrow.Record) (arrow.Record, error)
type SchemaTransformationFn = func(*arrow.Schema) (*arrow.Schema, error)

type Transformer struct {
	matcher  *tablematcher.TableMatcher
	fn       TransformationFn
	schemaFn SchemaTransformationFn
}

func NewFromSpec(sp spec.Spec) (*Transformer, error) {
	tr := &Transformer{matcher: tablematcher.New(sp.Tables)}

	tr.fn = FlattenJSONFields()
	tr.schemaFn = transformSchema(tr.fn)

	return tr, nil
}

func (tr *Transformer) Transform(record arrow.Record) (arrow.Record, error) {
	// Passthrough if the record's table is not a match to any of the spec's
	// tablepatterns, but error if the record doesn't have table metadata.
	isMatch, err := tr.matcher.IsSchemasTableMatch(record.Schema())
	if err != nil {
		return nil, err
	}
	if !isMatch {
		return record, nil
	}
	// Apply the specific transformation kind
	return tr.fn(record)
}

func (tr *Transformer) TransformSchema(schema *arrow.Schema) (*arrow.Schema, error) {
	// Passthrough if the record's table is not a match to any of the spec's
	// tablepatterns, but error if the record doesn't have table metadata.
	isMatch, err := tr.matcher.IsSchemasTableMatch(schema)
	if err != nil {
		return nil, err
	}
	if !isMatch {
		return schema, nil
	}

	if tr.schemaFn == nil {
		return schema, nil
	}
	return tr.schemaFn(schema)
}

func FlattenJSONFields() TransformationFn {
	return func(record arrow.Record) (arrow.Record, error) {
		record, err := recordupdater.New(record).FlattenJSONFields()
		if err != nil {
			return nil, err
		}
		return record, nil
	}
}

func transformSchema(tf TransformationFn) SchemaTransformationFn {
	return func(schema *arrow.Schema) (*arrow.Schema, error) {
		newRecord, err := tf(makeEmptyRecord(schema))
		if err != nil {
			return nil, err
		}
		return newRecord.Schema(), nil
	}
}

func makeEmptyRecord(s *arrow.Schema) arrow.Record {
	cols := []arrow.Array{}
	for _, field := range s.Fields() {
		cols = append(cols, array.NewBuilder(memory.DefaultAllocator, field.Type).NewArray())
	}
	md := s.Metadata()
	return array.NewRecord(arrow.NewSchema(s.Fields(), &md), cols, 0)
}
