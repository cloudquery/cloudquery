package transformer

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var transformTestCases = []struct {
	name               string
	transformer        func() *RecordTransformer
	originalSchema     *arrow.Schema
	originalJSONRecord []byte
	expectedSchema     *arrow.Schema
	expectedJSONRecord []byte
}{
	{
		name: "no_transformation",
		transformer: func() *RecordTransformer {
			return NewRecordTransformer()
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		expectedJSONRecord: []byte(`{"id": 1}`),
	},
	{
		name: "add_source",
		transformer: func() *RecordTransformer {
			return NewRecordTransformer(WithSourceNameColumn("test"))
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "_cq_source_name", Type: arrow.BinaryTypes.String, Nullable: true},
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		expectedJSONRecord: []byte(`{"_cq_source_name": "test","id": 1}`),
	},
	{
		name: "add_sync_time",
		transformer: func() *RecordTransformer {
			t, err := time.Parse(time.RFC3339, "2023-06-21T17:54:44.488177Z")
			if err != nil {
				panic(err)
			}
			return NewRecordTransformer(WithSyncTimeColumn(t))
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "_cq_sync_time", Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: true},
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		expectedJSONRecord: []byte(`{"_cq_sync_time": "2023-06-21 17:54:44.488177","id": 1}`),
	},
	{
		name: "add_source_and_sync_time",
		transformer: func() *RecordTransformer {
			t, err := time.Parse(time.RFC3339, "2023-06-21T17:54:44.488177Z")
			if err != nil {
				panic(err)
			}
			return NewRecordTransformer(WithSyncTimeColumn(t), WithSourceNameColumn("test"))
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "_cq_sync_time", Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: true},
			{Name: "_cq_source_name", Type: arrow.BinaryTypes.String, Nullable: true},
			{Name: "id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		expectedJSONRecord: []byte(`{"_cq_sync_time": "2023-06-21 17:54:44.488177","_cq_source_name": "test","id": 1}`),
	},
	{
		name: "use_cq_id_primary_key_with_remove_pks",
		transformer: func() *RecordTransformer {
			return NewRecordTransformer(WithRemovePKs(), WithCQIDPrimaryKey())
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{schema.MetadataPrimaryKey: "true"})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{schema.MetadataPrimaryKey: "true"})},
		}, nil),
		expectedJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
	},
	{
		name: "use_cq_id_primary_key_with_remove_unique",
		transformer: func() *RecordTransformer {
			return NewRecordTransformer(WithRemovePKs(), WithCQIDPrimaryKey(), WithRemoveUniqueConstraints())
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{schema.MetadataPrimaryKey: "true", schema.MetadataUnique: "true"})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{schema.MetadataPrimaryKey: "true"})},
		}, nil),
		expectedJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
	},
	{
		name: "use_with_remove_unique",
		transformer: func() *RecordTransformer {
			return NewRecordTransformer(WithRemovePKs(), WithRemoveUniqueConstraints())
		},
		originalSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{schema.MetadataPrimaryKey: "true", schema.MetadataUnique: "true"})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		originalJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
		expectedSchema: arrow.NewSchema([]arrow.Field{
			{Name: "id", Type: arrow.PrimitiveTypes.Int64, Metadata: arrow.MetadataFrom(map[string]string{})},
			{Name: "_cq_id", Type: arrow.PrimitiveTypes.Int64},
		}, nil),
		expectedJSONRecord: []byte(`{"id": 1, "_cq_id": 2}`),
	},
}

func TestRecord(t *testing.T) {
	for _, tc := range transformTestCases {
		t.Run(tc.name, func(t *testing.T) {
			bldr := array.NewRecordBuilder(memory.DefaultAllocator, tc.originalSchema)
			if err := bldr.UnmarshalJSON(tc.originalJSONRecord); err != nil {
				t.Fatal(err)
			}
			record := bldr.NewRecord()
			transformedRecord := tc.transformer().Transform(record)
			if transformedRecord.Schema().String() != tc.expectedSchema.String() {
				t.Fatalf("expected schema\n%v, got\n%v", tc.expectedSchema, transformedRecord.Schema())
			}
			bldr = array.NewRecordBuilder(memory.DefaultAllocator, transformedRecord.Schema())
			if err := bldr.UnmarshalJSON(tc.expectedJSONRecord); err != nil {
				t.Fatal(err)
			}
			expectedRecord := bldr.NewRecord()
			if !array.RecordEqual(expectedRecord, transformedRecord) {
				b, err := json.Marshal(transformedRecord)
				if err != nil {
					t.Fatal(err)
				}
				t.Logf("expected record %v, got %v", string(tc.expectedJSONRecord), string(b))
				t.Fatalf("expected record %v, got %v", expectedRecord, transformedRecord)
			}
		})
	}
}
