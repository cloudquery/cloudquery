package transformer

import (
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/apache/arrow/go/v16/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	nRows := int(record.NumRows())

	cols := make([]arrow.Array, 0, len(sc.Fields())+t.internalColumns)
	if t.withSyncTime && !sc.HasField(cqSyncTime) {
		ts, _ := arrow.TimestampFromTime(t.syncTime, arrow.Microsecond)
		syncTimeBldr := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
		syncTimeBldr.Reserve(nRows)
		for i := 0; i < nRows; i++ {
			syncTimeBldr.Append(ts)
		}
		cols = append(cols, syncTimeBldr.NewArray())
	}
	if t.withSourceName && !sc.HasField(cqSourceName) {
		sourceBldr := array.NewStringBuilder(memory.DefaultAllocator)
		sourceBldr.Reserve(nRows)
		for i := 0; i < nRows; i++ {
			sourceBldr.Append(t.sourceName)
		}
		cols = append(cols, sourceBldr.NewArray())
	}
	if t.withSyncGroupID && !sc.HasField(cqSyncGroupId) {
		syncGroupIdBldr := array.NewStringBuilder(memory.DefaultAllocator)
		syncGroupIdBldr.Reserve(nRows)
		for i := 0; i < nRows; i++ {
			syncGroupIdBldr.Append(t.syncGroupId)
		}
		cols = append(cols, syncGroupIdBldr.NewArray())
	}

	cols = append(cols, record.Columns()...)

	return array.NewRecord(newSchema, cols, int64(nRows))
}
