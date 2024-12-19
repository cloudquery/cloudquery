package transformers

import (
	"bytes"
	"testing"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/transformer/jsonflattener/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestTransform(t *testing.T) {
	tests := []struct {
		name     string
		spec     spec.Spec
		record   arrow.Record
		validate func(t *testing.T, record arrow.Record)
	}{
		{
			name: "FlattenJSONFields",
			spec: spec.Spec{
				Tables: []string{"*"},
			},
			record: createTestRecord(t),
			validate: func(t *testing.T, record arrow.Record) {
				require.Equal(t, int64(4), record.NumCols())
				require.Equal(t, int64(1), record.NumRows())
				requireAllColsLenMatchRecordsLen(t, record)
				require.Equal(t, "col1__key_a", record.ColumnName(1))
				require.Equal(t, "value", record.Column(1).(*array.String).Value(0))
				require.Equal(t, "col1__key_b", record.ColumnName(2))
				require.Equal(t, int64(2), record.Column(2).(*array.Int64).Value(0))
				require.Equal(t, "col1__key_c", record.ColumnName(3))
				require.Equal(t, true, record.Column(3).(*array.Boolean).Value(0))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformer, err := NewFromSpec(zerolog.Nop(), tt.spec)
			require.NoError(t, err, "NewFromSpec() should not return an error")

			transformedRecord, err := transformer.Transform(tt.record)
			require.NoError(t, err, "Transform() should not return an error")

			requireAllColsLenMatchRecordsLen(t, transformedRecord)
			tt.validate(t, transformedRecord)
		})
	}
}

const sampleRowContent = `{"key_a": "value", "key_b": 2, "key_c": true}`

func createTestRecord(t *testing.T) arrow.Record {
	return array.NewRecord(createTestSchema(t), []arrow.Array{createArray(t, []byte(sampleRowContent))}, 1)
}

func requireAllColsLenMatchRecordsLen(t *testing.T, record arrow.Record) {
	for i := 0; i < int(record.NumCols()); i++ {
		require.Equal(t, int(record.NumRows()), record.Column(i).Len(), "Expected length of %d for column %d", record.NumRows(), i)
	}
}

func createTestSchema(t *testing.T) *arrow.Schema {
	md := arrow.NewMetadata([]string{schema.MetadataTableName}, []string{"testTable"})
	arr := createArray(t, []byte(sampleRowContent))
	fieldMD := map[string]string{
		schema.MetadataTypeSchema: `{"key_a": "utf8", "key_b": "int64", "key_c": "bool"}`,
	}
	return arrow.NewSchema(
		[]arrow.Field{
			{Name: "col1", Type: arr.DataType(), Nullable: true, Metadata: arrow.MetadataFrom(fieldMD)},
		},
		&md,
	)
}

func createArray(t *testing.T, bs []byte) arrow.Array {
	b := types.NewJSONBuilder(memory.NewGoAllocator())
	defer b.Release()
	dec := json.NewDecoder(bytes.NewReader(bs))
	err := b.UnmarshalOne(dec)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return b.NewArray()
}
