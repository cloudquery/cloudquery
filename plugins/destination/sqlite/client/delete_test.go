package client

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/apache/arrow-go/v18/arrow/array"
	"github.com/apache/arrow-go/v18/arrow/memory"
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
	ch := make(chan arrow.RecordBatch)
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
	s := &Spec{ConnectionString: ":memory:"}
	b, err := json.Marshal(s)
	r.NoError(err)
	c, err := New(ctx, zerolog.Nop(), b, plugin.NewClientOptions{})
	r.NoError(err)
	return c.(*Client)
}

func valueToArrowRecord(tableName string, value string) arrow.RecordBatch {
	bldrDeleteMatch := array.NewRecordBuilder(memory.DefaultAllocator, (&schema.Table{
		Name: tableName,
		Columns: schema.ColumnList{
			schema.Column{Name: "id", Type: arrow.BinaryTypes.String},
		},
	}).ToArrowSchema())
	bldrDeleteMatch.Field(0).(*array.StringBuilder).Append(value)
	deleteValue := bldrDeleteMatch.NewRecordBatch()
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
	writeInserts := make(message.WriteInserts, 0, len(values))
	for _, insertValue := range values {
		bldr := array.NewRecordBuilder(memory.DefaultAllocator, table.ToArrowSchema())
		bldr.Field(0).(*array.StringBuilder).Append(insertValue)
		bldr.Field(1).(*array.StringBuilder).Append(sourceName)
		bldr.Field(2).(*array.TimestampBuilder).AppendTime(time.Now())
		record := bldr.NewRecordBatch()
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

func TestIdentifier(t *testing.T) {
	testCases := map[string]string{
		"plain":        `"plain"`,
		`with"quote`:   `"with""quote"`,
		`a"; drop; --`: `"a""; drop; --"`,
		`""`:           `""""""`,
		"":             `""`,
	}
	for input, expected := range testCases {
		require.Equal(t, expected, identifier(input))
	}
}

func TestIdentifiersWithQuotesAndSemicolons(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()
	client := withPluginClient(ctx, r)

	const idColumn = `id"; --`
	table := createTestTable()
	table.Name = `quoted" (x); CREATE TABLE injected (y); --`
	table.Columns[0].Name = idColumn
	r.NoError(client.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))
	r.Equal([]string{table.Name}, tableNames(ctx, r, client))

	r.NoError(client.WriteTableBatch(ctx, "", createInsertMessages([]string{"foo", "bar"}, table)))

	table.Columns = append(table.Columns, schema.Column{Name: `added"; --`, Type: arrow.BinaryTypes.String})
	r.NoError(client.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))
	count, err := countAllRows(ctx, client, table)
	r.NoError(err)
	r.EqualValues(2, count)

	deletes := createDeleteMessages(false, table, []string{"foo"})
	deletes[0].WhereClause[0].Predicates[0].Column = idColumn
	r.NoError(client.DeleteRecord(ctx, deletes))
	count, err = countAllRows(ctx, client, table)
	r.NoError(err)
	r.EqualValues(1, count)

	r.NoError(client.DeleteStale(ctx, message.WriteDeleteStales{{
		TableName:  table.Name,
		SourceName: "source-test",
		SyncTime:   time.Now().Add(time.Hour),
	}}))
	count, err = countAllRows(ctx, client, table)
	r.NoError(err)
	r.EqualValues(0, count)
	r.Equal([]string{table.Name}, tableNames(ctx, r, client))
}

func TestDeleteRejectsUnknownGroupingType(t *testing.T) {
	r := require.New(t)
	ctx := context.Background()
	client := withPluginClient(ctx, r)

	table := createTestTable()
	r.NoError(client.MigrateTables(ctx, message.WriteMigrateTables{{Table: table}}))
	r.NoError(client.WriteTableBatch(ctx, "", createInsertMessages([]string{"foo", "bar"}, table)))

	for _, groupingType := range []string{"OR 1 = 1 OR", "XOR", ""} {
		deletes := createDeleteMessages(false, table, []string{"missing"})
		group := &deletes[0].WhereClause[0]
		group.GroupingType = groupingType
		group.Predicates = append(group.Predicates, group.Predicates[0])
		err := client.DeleteRecord(ctx, deletes)
		r.ErrorContains(err, "unsupported grouping type")
	}
	count, err := countAllRows(ctx, client, table)
	r.NoError(err)
	r.EqualValues(2, count)

	for _, groupingType := range []string{"AND", "OR", "or"} {
		sql, err := generateDelete(message.DeleteRecord{
			TableName: table.Name,
			WhereClause: message.PredicateGroups{{
				GroupingType: groupingType,
				Predicates:   []message.Predicate{{Column: "a"}, {Column: "b"}},
			}},
		})
		r.NoError(err)
		r.Contains(sql, `("a" = ? `+strings.ToUpper(groupingType)+` "b" = ?)`)
	}
}

func tableNames(ctx context.Context, r *require.Assertions, client *Client) []string {
	rows, err := client.db.QueryContext(ctx, "SELECT name FROM sqlite_master WHERE type = 'table'")
	r.NoError(err)
	defer rows.Close()
	var names []string
	for rows.Next() {
		var name string
		r.NoError(rows.Scan(&name))
		names = append(names, name)
	}
	r.NoError(rows.Err())
	return names
}
