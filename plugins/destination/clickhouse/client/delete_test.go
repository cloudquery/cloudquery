package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v7/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	testCases := []struct {
		name          string
		insertValues  []string
		deleteValues  []string
		deleteAll     bool
		expectedCount int
	}{
		{
			name:          "delete single record",
			insertValues:  []string{"foo", "bar"},
			deleteValues:  []string{"foo"},
			expectedCount: 1,
		},
		{
			name:          "delete both records",
			insertValues:  []string{"foo", "bar"},
			deleteValues:  []string{"foo", "bar"},
			expectedCount: 0,
		},
		{
			name:          "delete none",
			insertValues:  []string{"foo"},
			deleteValues:  []string{"bar"},
			expectedCount: 1,
		},
		{
			name:          "delete all records",
			insertValues:  []string{"foo", "bar"},
			deleteAll:     true,
			expectedCount: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)
			ctx := context.Background()
			client := withPluginClient(ctx, r)

			table := createTestTable()
			r.NoError(client.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))

			writeInserts := createInsertMessages(tc.insertValues, table)
			r.NoError(client.WriteTableBatch(ctx, "", writeInserts))

			writeDeletes := createDeleteMessages(tc.deleteAll, table, tc.deleteValues)
			r.NoError(client.DeleteRecord(ctx, writeDeletes))

			count, err := countAllRows(ctx, client, table)
			r.NoError(err)
			r.EqualValues(tc.expectedCount, count, "unexpected amount of items after delete with match")
		})
	}
}

func countAllRows(ctx context.Context, client *Client, table *schema.Table) (int64, error) {
	var err error
	ch := make(chan arrow.Record)
	go func() {
		defer close(ch)
		err = client.Read(ctx, table, ch)
	}()
	count := int64(0)
	for record := range ch {
		count += record.NumRows()
	}
	return count, err
}

func withPluginClient(ctx context.Context, r *require.Assertions) *Client {
	s := &spec.Spec{ConnectionString: getTestConnection()}
	b, err := json.Marshal(s)
	r.NoError(err)
	c, err := New(ctx, zerolog.Nop(), b, plugin.NewClientOptions{})
	r.NoError(err)
	return c.(*Client)
}

func valueToArrowRecord(tableName string, value string) arrow.Record {
	bldrDeleteMatch := array.NewRecordBuilder(memory.DefaultAllocator, (&schema.Table{
		Name: tableName,
		Columns: schema.ColumnList{
			schema.Column{Name: "id", Type: arrow.BinaryTypes.String},
		},
	}).ToArrowSchema())
	bldrDeleteMatch.Field(0).(*array.StringBuilder).Append(value)
	deleteValue := bldrDeleteMatch.NewRecord()
	return deleteValue
}

func createDeleteMessages(deleteAll bool, table *schema.Table, deleteValues []string) message.WriteDeleteRecords {
	writeDeletes := message.WriteDeleteRecords{}

	if deleteAll {
		msg := message.WriteDeleteRecord{
			DeleteRecord: message.DeleteRecord{
				TableName: table.Name,
			},
		}
		return append(writeDeletes, &msg)
	}
	for _, deleteValue := range deleteValues {
		msg := message.WriteDeleteRecord{
			DeleteRecord: message.DeleteRecord{
				TableName: table.Name,
				WhereClause: message.PredicateGroups{
					{
						GroupingType: "AND",
						Predicates: []message.Predicate{
							{
								Operator: "eq",
								Column:   "id",
								Record:   valueToArrowRecord(table.Name, deleteValue),
							},
						},
					},
				},
			},
		}
		writeDeletes = append(writeDeletes, &msg)
	}
	return writeDeletes
}

func createInsertMessages(values []string, table *schema.Table) message.WriteInserts {
	const sourceName = "source-test"
	writeInserts := message.WriteInserts{}
	for _, insertValue := range values {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		bldr.Field(0).(*array.StringBuilder).Append(insertValue)
		bldr.Field(1).(*array.StringBuilder).Append(sourceName)
		bldr.Field(2).(*array.TimestampBuilder).AppendTime(time.Now())
		record := bldr.NewRecord()
		writeInserts = append(writeInserts, &message.WriteInsert{Record: record})
	}
	return writeInserts
}

func createTestTable() *schema.Table {
	tableName := fmt.Sprintf("cq_delete_test_%d_%04d", time.Now().UnixNano(), rand.Intn(1000))
	table := &schema.Table{
		Name: tableName,
		Columns: schema.ColumnList{
			schema.Column{Name: "id", Type: arrow.BinaryTypes.String, PrimaryKey: true, NotNull: true},
			schema.CqSourceNameColumn,
			schema.CqSyncTimeColumn,
		},
	}
	return table
}
