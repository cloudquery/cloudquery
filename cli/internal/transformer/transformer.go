package transformer

import (
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/apache/arrow/go/v13/arrow/memory"
)

const (
	cqSyncTime     = "_cq_sync_time"
	cqSourceName   = "_cq_source_name"
	cqIDColumnName = "_cq_id"
)

type RecordTransformer struct {
	sourceName      string
	withSourceName  bool
	syncTime        time.Time
	withSyncTime    bool
	internalColumns int
	removePks       bool
	cqIDPrimaryKey  bool
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

func WithRemovePKs() RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.removePks = true
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
	if t.withSyncTime {
		fields = append(fields, arrow.Field{Name: cqSyncTime, Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: true})
	}
	if t.withSourceName {
		fields = append(fields, arrow.Field{Name: cqSourceName, Type: arrow.BinaryTypes.String, Nullable: true})
	}
	for _, field := range sc.Fields() {
		mdMap := field.Metadata.ToMap()
		if _, ok := mdMap["cq:extension:primary_key"]; ok && t.removePks {
			delete(mdMap, "cq:extension:primary_key")
		}
		if field.Name == cqIDColumnName && t.cqIDPrimaryKey {
			mdMap["cq:extension:primary_key"] = "true"
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
	if t.withSyncTime {
		syncTimeBldr := array.NewTimestampBuilder(memory.DefaultAllocator, &arrow.TimestampType{Unit: arrow.Microsecond, TimeZone: "UTC"})
		for i := 0; i < nRows; i++ {
			syncTimeBldr.AppendTime(t.syncTime)
		}
		cols = append(cols, syncTimeBldr.NewArray())
	}
	if t.withSourceName {
		sourceBldr := array.NewStringBuilder(memory.DefaultAllocator)
		for i := 0; i < nRows; i++ {
			sourceBldr.Append(t.sourceName)
		}
		cols = append(cols, sourceBldr.NewArray())
	}

	for i := range sc.Fields() {
		cols = append(cols, record.Column(i))
	}

	return array.NewRecord(newSchema, cols, int64(nRows))
}
