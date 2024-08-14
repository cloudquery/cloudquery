package transformers

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/apache/arrow/go/v17/arrow/array"
	"github.com/apache/arrow/go/v17/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/sql/client/spec"
	"github.com/cloudquery/cloudquery/plugins/transformer/sql/client/sqlrunner"
	"github.com/cloudquery/cloudquery/plugins/transformer/sql/client/tablematcher"
)

var globalSQLRunner *sqlrunner.DuckDBSQLRunner

type TransformationFn = func(arrow.Record) ([]arrow.Record, error)
type SchemaTransformationFn = func(*arrow.Schema) (*arrow.Schema, error)

type Transformer struct {
	matcher  *tablematcher.TableMatcher
	fn       TransformationFn
	schemaFn SchemaTransformationFn
}

func NewFromSpec(ctx context.Context, sp spec.TransformationSpec) (*Transformer, error) {
	tr := &Transformer{matcher: tablematcher.New(sp.Tables)}

	switch sp.Kind {
	case spec.KindSQL:
		err := registerSQLRunner(ctx)
		if err != nil {
			return nil, err
		}
		tr.fn = RunSQL(sp.SQL)
	default:
		return nil, fmt.Errorf("unknown transformation kind: %s", sp.Kind)
	}

	tr.schemaFn = transformSchema(tr.fn)

	return tr, nil
}

func registerSQLRunner(ctx context.Context) error {
	if globalSQLRunner != nil {
		return nil
	}
	var err error
	globalSQLRunner, err = sqlrunner.New(ctx)
	return err
}

func (tr *Transformer) Transform(record arrow.Record) ([]arrow.Record, error) {
	// Passthrough if the record's table is not a match to any of the spec's
	// tablepatterns, but error if the record doesn't have table metadata.
	isMatch, err := tr.matcher.IsSchemasTableMatch(record.Schema())
	if err != nil {
		return nil, err
	}
	if !isMatch {
		return []arrow.Record{record}, nil
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

func RunSQL(sql string) TransformationFn {
	return func(record arrow.Record) ([]arrow.Record, error) {
		return globalSQLRunner.RunSQLOnRecord(record, sql)
	}
}

func transformSchema(tf TransformationFn) SchemaTransformationFn {
	return func(schema *arrow.Schema) (*arrow.Schema, error) {
		newRecords, err := tf(makeRecordWithZeroValue(schema))
		if err != nil {
			return nil, err
		}
		if len(newRecords) == 0 {
			return nil, fmt.Errorf("transformation of MigrateTable record resulted in 0 records; can't migrate schema")
		}
		return newRecords[0].Schema(), nil
	}
}

func makeRecordWithZeroValue(s *arrow.Schema) arrow.Record {
	cols := []arrow.Array{}
	for _, field := range s.Fields() {
		bldr := array.NewBuilder(memory.DefaultAllocator, field.Type)
		bldr.AppendEmptyValue()
		cols = append(cols, bldr.NewArray())
	}
	md := s.Metadata()
	return array.NewRecord(arrow.NewSchema(s.Fields(), &md), cols, 1)
}
