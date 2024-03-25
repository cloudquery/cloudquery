package client

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/apache/arrow/go/v15/arrow/array"
	"github.com/apache/arrow/go/v15/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func (c *Client) addObjectsSyncedToSummary(record arrow.Record) arrow.Record {
	sc := record.Schema()
	newSchema := transformSchema(sc)
	nRows := int(record.NumRows())

	cols := make([]arrow.Array, 0, len(newSchema.Fields())+1)

	objKeyBuilder := types.NewJSONBuilder(array.NewExtensionBuilder(memory.DefaultAllocator, types.NewJSONType()))
	for i := 0; i < nRows; i++ {
		objKeyBuilder.Append(c.objectKeys)
	}
	cols = append(cols, objKeyBuilder.NewArray())

	for i := range sc.Fields() {
		cols = append(cols, record.Column(i))
	}

	return array.NewRecord(newSchema, cols, int64(nRows))
}

func transformSchema(sc *arrow.Schema) *arrow.Schema {
	fields := make([]arrow.Field, 0, len(sc.Fields())+1)
	fields = append(fields, arrow.Field{Name: "objects_synced", Type: &types.JSONType{}, Nullable: true})
	for _, field := range sc.Fields() {
		mdMap := field.Metadata.ToMap()

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
