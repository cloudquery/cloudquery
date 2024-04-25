package transformer

import (
	"time"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/exp/constraints"
)

const (
	cqSyncTime     = "_cq_sync_time"
	cqSourceName   = "_cq_source_name"
	cqIDColumnName = "_cq_id"
	cqSyncGroupId  = "_cq_sync_group_id"
)

type RecordTransformer struct {
	sourceName              string
	withSourceName          bool
	syncTime                time.Time
	withSyncTime            bool
	internalColumns         int
	removePks               bool
	removeUniqueConstraints bool
	cqIDPrimaryKey          bool
	withSyncGroupID         bool
	syncGroupId             string
}

type RecordTransformerOption func(*RecordTransformer)

func WithSourceNameColumn(name string) RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.sourceName = name
		transformer.withSourceName = true
		transformer.internalColumns++
	}
}

func WithSyncTimeColumn(t time.Time) RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.syncTime = t
		transformer.withSyncTime = true
		transformer.internalColumns++
	}
}

func WithSyncGroupIdColumn(syncGroupId string) RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.withSyncGroupID = true
		transformer.syncGroupId = syncGroupId
		transformer.internalColumns++
	}
}

func WithRemovePKs() RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.removePks = true
	}
}

func WithRemoveUniqueConstraints() RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.removeUniqueConstraints = true
	}
}

func WithCQIDPrimaryKey() RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.cqIDPrimaryKey = true
	}
}

func NewRecordTransformer(opts ...RecordTransformerOption) *RecordTransformer {
	t := &RecordTransformer{}
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *RecordTransformer) TransformSchema(sc *arrow.Schema) *arrow.Schema {
	fields := make([]arrow.Field, 0, len(sc.Fields())+t.internalColumns)
	if t.withSyncTime && !sc.HasField(cqSyncTime) {
		fields = append(fields, arrow.Field{Name: cqSyncTime, Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: true})
	}
	if t.withSourceName && !sc.HasField(cqSourceName) {
		fields = append(fields, arrow.Field{Name: cqSourceName, Type: arrow.BinaryTypes.String, Nullable: true})
	}
	if t.withSyncGroupID && !sc.HasField(cqSyncGroupId) {
		fields = append(fields, arrow.Field{
			Name: cqSyncGroupId,
			Type: arrow.BinaryTypes.String,
			Metadata: arrow.NewMetadata(
				[]string{schema.MetadataPrimaryKey},
				[]string{schema.MetadataTrue},
			)})
	}
	for _, field := range sc.Fields() {
		mdMap := field.Metadata.ToMap()

		if _, ok := mdMap[schema.MetadataUnique]; ok && t.removeUniqueConstraints {
			delete(mdMap, schema.MetadataUnique)
		}

		if _, ok := mdMap[schema.MetadataPrimaryKey]; ok && t.removePks {
			delete(mdMap, schema.MetadataPrimaryKey)
		}
		if field.Name == cqIDColumnName && t.cqIDPrimaryKey {
			mdMap[schema.MetadataPrimaryKey] = schema.MetadataTrue
		}

		newMd := arrow.MetadataFrom(mdMap)

		fields = append(fields, arrow.Field{
			Name:     field.Name,
			Type:     field.Type,
			Nullable: field.Nullable,
			Metadata: newMd,
		})
	}
	scMd := sc.Metadata()
	return arrow.NewSchema(fields, &scMd)
}

func (t *RecordTransformer) Transform(record arrow.Record) arrow.Record {
	sc := record.Schema()
	newSchema := t.TransformSchema(sc)
	nRows := record.NumRows()

	cols := make([]arrow.Array, 0, len(sc.Fields())+t.internalColumns) // alloc together with the proper capacity
	if t.withSyncTime && !sc.HasField(cqSyncTime) {
		ts, _ := arrow.TimestampFromTime(t.syncTime, arrow.Microsecond)
		builder := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
		builder.AppendValues(repeat(ts, nRows), nil)
		cols = append(cols, builder.NewArray())
	}
	if t.withSourceName && !sc.HasField(cqSourceName) {
		builder := array.NewStringBuilder(memory.DefaultAllocator)
		builder.AppendStringValues(repeat(t.sourceName, nRows), nil)
		cols = append(cols, builder.NewArray())
	}
	if t.withSyncGroupID && !sc.HasField(cqSyncGroupId) {
		builder := array.NewStringBuilder(memory.DefaultAllocator)
		builder.AppendStringValues(repeat(t.syncGroupId, nRows), nil)
		cols = append(cols, builder.NewArray())
	}

	cols = cols[:len(sc.Fields())+t.internalColumns] // resize back as we have the capacity
	copy(cols[t.internalColumns:], record.Columns())

	return array.NewRecord(newSchema, cols, nRows)
}

func repeat[A any, L constraints.Integer](val A, n L) []A {
	res := make([]A, n)
	for i := L(0); i < n; i++ {
		res[i] = val
	}
	return res
}
