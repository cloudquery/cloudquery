package client

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func transformArr(arr arrow.Array) []any {
	dbArr := make([]any, arr.Len())
	for i := 0; i < arr.Len(); i++ {
		if arr.IsNull(i) {
			continue
		}
		switch a := arr.(type) {
		case *array.Boolean:
			dbArr[i] = a.Value(i)
		case *array.Int8:
			dbArr[i] = a.Value(i)
		case *array.Int16:
			dbArr[i] = a.Value(i)
		case *array.Int32:
			dbArr[i] = a.Value(i)
		case *array.Int64:
			dbArr[i] = a.Value(i)
		case *array.Uint8:
			dbArr[i] = a.Value(i)
		case *array.Uint16:
			dbArr[i] = a.Value(i)
		case *array.Uint32:
			dbArr[i] = a.Value(i)
		case *array.Uint64:
			val := a.Value(i)
			var custom = CustomUnit64(val)
			dbArr[i] = custom
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
			dbArr[i] = a.Value(i).ToTime(a.DataType().(*arrow.TimestampType).Unit)
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

func (*Client) transformRecord(table *schema.Table, record arrow.Record) []any {
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
			documents[l].(bson.M)[table.Columns[i].Name] = transformed[l]
		}
	}
	return documents
}

func (c *Client) transformRecords(table *schema.Table, records []arrow.Record) []any {
	documents := make([]any, 0, len(records))
	for _, r := range records {
		docs := c.transformRecord(table, r)
		documents = append(documents, docs...)
	}
	return documents
}

func (c *Client) appendTableBatch(ctx context.Context, table *schema.Table, documents []any) error {
	tableName := table.Name
	if _, err := c.client.Database(c.spec.Database).Collection(tableName).InsertMany(ctx, documents); err != nil {
		return err
	}
	return nil
}

func (c *Client) overwriteTableBatch(ctx context.Context, table *schema.Table, documents []any) error {
	tableName := table.Name
	operations := make([]mongo.WriteModel, len(documents))
	pks := table.PrimaryKeys()
	for i, document := range documents {
		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		filter := make(bson.M, len(pks))
		for _, name := range pks {
			filter[name] = document.(bson.M)[name]
		}
		operation.SetFilter(filter)
		update := make(bson.M, len(table.Columns))
		for _, col := range table.Columns {
			update[col.Name] = document.(bson.M)[col.Name]
		}
		operation.SetUpdate(bson.M{"$set": update})
		operations[i] = operation
	}
	if _, err := c.client.Database(c.spec.Database).Collection(tableName).BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}

func (c *Client) WriteTableBatch(ctx context.Context, tableName string, msgs message.WriteInserts) error {
	if len(msgs) == 0 {
		return nil
	}
	table, err := schema.NewTableFromArrowSchema(msgs[0].Record.Schema())
	if err != nil {
		return err
	}
	records := make([]arrow.Record, len(msgs))
	for i, msg := range msgs {
		records[i] = msg.Record
	}
	documents := c.transformRecords(table, records)
	if len(table.PrimaryKeys()) > 0 {
		return c.overwriteTableBatch(ctx, table, documents)
	}
	return c.appendTableBatch(ctx, table, documents)
}

func (c *Client) Write(ctx context.Context, msgs <-chan message.WriteMessage) error {
	if err := c.writer.Write(ctx, msgs); err != nil {
		return err
	}
	if err := c.writer.Flush(ctx); err != nil {
		return fmt.Errorf("failed to flush: %w", err)
	}
	return nil
}
