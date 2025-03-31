package transformer

import (
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
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
	cqColumnsNotNull        bool
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

func WithCQColumnsNotNull() RecordTransformerOption {
	return func(transformer *RecordTransformer) {
		transformer.cqColumnsNotNull = true
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
		fields = append(fields, arrow.Field{Name: cqSyncTime, Type: arrow.FixedWidthTypes.Timestamp_us, Nullable: !t.cqColumnsNotNull})
	}
	if t.withSourceName && !sc.HasField(cqSourceName) {
		fields = append(fields, arrow.Field{Name: cqSourceName, Type: arrow.BinaryTypes.String, Nullable: !t.cqColumnsNotNull})
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
	fields = append(fields, sc.Fields()...)

	transformedFields := make([]arrow.Field, len(fields))
	for i, field := range fields {
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

		transformedFields[i] = arrow.Field{
			Name:     field.Name,
			Type:     field.Type,
			Nullable: field.Nullable,
			Metadata: arrow.MetadataFrom(mdMap),
		}
	}
	scMd := sc.Metadata()
	return arrow.NewSchema(transformedFields, &scMd)
}

func (t *RecordTransformer) Transform(record arrow.Record) arrow.Record {
	sc := record.Schema()
	newSchema := t.TransformSchema(sc)
	nRows := int(record.NumRows())

	cols := make([]arrow.Array, 0, len(sc.Fields())+t.internalColumns)
	if t.withSyncTime && (!sc.HasField(cqSyncTime) || env.IsCloud()) {
		syncTimeArray, _ := schema.TimestampArrayFromTime(t.syncTime, arrow.Microsecond, "UTC", nRows)
		if !sc.HasField(cqSyncTime) {
			cols = append(cols, syncTimeArray)
		} else {
			record, _ = schema.ReplaceFieldInRecord(record, cqSyncTime, syncTimeArray)
		}
	}
	if t.withSourceName && (!sc.HasField(cqSourceName) || env.IsCloud()) {
		sourceNameArray := schema.StringArrayFromValue(t.sourceName, nRows)
		if !sc.HasField(cqSourceName) {
			cols = append(cols, sourceNameArray)
		} else {
			record, _ = schema.ReplaceFieldInRecord(record, cqSourceName, sourceNameArray)
		}
	}
	if t.withSyncGroupID && (!sc.HasField(cqSyncGroupId) || env.IsCloud()) {
		syncGroupIdArray := schema.StringArrayFromValue(t.syncGroupId, nRows)
		if !sc.HasField(cqSyncGroupId) {
			cols = append(cols, syncGroupIdArray)
		} else {
			record, _ = schema.ReplaceFieldInRecord(record, cqSyncGroupId, syncGroupIdArray)
		}
	}

	cols = append(cols, record.Columns()...)

	return array.NewRecord(newSchema, cols, int64(nRows))
}
