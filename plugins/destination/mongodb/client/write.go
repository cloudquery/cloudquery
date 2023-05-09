package client

import (
	"context"
	"encoding/json"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/apache/arrow/go/v13/arrow/array"
	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func transformArr(arr arrow.Array) []any {
	dbArr := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) || !arr.IsValid(i) {
			dbArr[i] = nil
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			dbArr[i] = a.Value(i)
		case *array.Int16:
			dbArr[i] = a.Value(i)
		case *array.Int32:
			dbArr[i] = a.Value(i)
		case *array.Int64:
			dbArr[i] = a.Value(i)
		case *array.Float32:
			dbArr[i] = a.Value(i)
		case *array.Float64:
			dbArr[i] = a.Value(i)
		case *array.Binary:
			dbArr[i] = a.Value(i)
		case *array.LargeBinary:
			dbArr[i] = a.Value(i)
		case *array.String:
			dbArr[i] = a.Value(i)
		case *array.LargeString:
			dbArr[i] = a.Value(i)
		case *array.Timestamp:
			dbArr[i] = a.Value(i).ToTime(arrow.Microsecond)
		case *types.JSONArray:
			var val any
			if err := json.Unmarshal([]byte(a.ValueStr(i)), &val); err != nil {
				panic(err)
			}
			dbArr[i] = val
		case *array.Struct:
			var val any
			if err := json.Unmarshal([]byte(a.ValueStr(i)), &val); err != nil {
				panic(err)
			}
			dbArr[i] = val
		case array.ListLike:
			start, end := a.ValueOffsets(i)
			nested := array.NewSlice(a.ListValues(), start, end)
			dbArr[i] = transformArr(nested)
		default:
			dbArr[i] = arr.ValueStr(i)
		}
	}

	return dbArr
}

func (*Client) transformRecord(table *arrow.Schema, record arrow.Record) []any {
	nc := int(record.NumCols())
	nr := int(record.NumRows())
	documents := make([]any, nr)
	for i := 0; i < nr; i++ {
		documents[i] = make(bson.M, nc)
	}

	for i := 0; i < nc; i++ {
		col := record.Column(i)
		transformed := transformArr(col)
		for l := 0; l < nr; l++ {
			documents[l].(bson.M)[table.Field(i).Name] = transformed[l]
		}
	}
	return documents
}

func (c *Client) transformRecords(table *arrow.Schema, records []arrow.Record) []any {
	documents := make([]any, 0, len(records))
	for _, r := range records {
		docs := c.transformRecord(table, r)
		documents = append(documents, docs...)
	}
	return documents
}

func (c *Client) appendTableBatch(ctx context.Context, table *arrow.Schema, docuemnts []any) error {
	tableName := schema.TableName(table)
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).InsertMany(ctx, docuemnts); err != nil {
		return err
	}
	return nil
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *arrow.Schema, documents []any) error {
	tableName := schema.TableName(table)
	operations := make([]mongo.WriteModel, len(documents))
	pks := schema.PrimaryKeyIndices(table)
	for i, document := range documents {
		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		filter := make(bson.M, len(pks))
		for _, pk := range pks {
			filter[table.Field(pk).Name] = document.(bson.M)[table.Field(pk).Name]
		}
		operation.SetFilter(filter)
		update := make(bson.M, len(table.Fields()))
		for _, col := range table.Fields() {
			update[col.Name] = document.(bson.M)[col.Name]
		}
		operation.SetUpdate(bson.M{"$set": update})
		operations[i] = operation
	}
	if _, err := c.client.Database(c.pluginSpec.Database).Collection(tableName).BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, table *arrow.Schema, resources []arrow.Record) error {
	documents := c.transformRecords(table, resources)

	pks := schema.PrimaryKeyIndices(table)
	if len(pks) == 0 {
		return c.appendTableBatch(ctx, table, documents)
	}
	switch c.spec.WriteMode {
	case specs.WriteModeAppend:
		return c.appendTableBatch(ctx, table, documents)
	case specs.WriteModeOverwrite, specs.WriteModeOverwriteDeleteStale:
		return c.overwriteTableBatch(ctx, table, documents)
	default:
		panic("unsupported write mode " + c.spec.WriteMode.String())
	}
}
