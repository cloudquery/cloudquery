package client

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v12/arrow"
	gremlingo "github.com/apache/tinkerpop/gremlin-go/v3/driver"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func (c *Client) Read(ctx context.Context, table *arrow.Schema, sourceName string, res chan<- arrow.Record) error {
	session, closer, err := c.newSession()
	if err != nil {
		return err
	}
	defer closer()

	tableName := schema.TableName(table)
	g := gremlingo.Traversal_().WithRemote(session).
		V().
		HasLabel(tableName).
		Has(schema.CqSourceNameColumn.Name, sourceName).
		Group().By(gremlingo.T.Id).
		By(AnonT.ValueMap())

	rs, err := g.GetResultSet()
	if err != nil {
		return fmt.Errorf("GetResultSet: %w", err)
	}
	defer rs.Close()

	//var records []arrow.Record
	for row := range rs.Channel() {
		m := row.Data.(map[any]any)
		for _, rowCols := range m {
			rowData := rowCols.(map[any]any)
			rec, err := reverseTransformer(table, rowData)
			if err != nil {
				return err
			}
			//records = append(records, rec)
			res <- rec
		}
	}

	//syncTimeIndex := table.FieldIndices(schema.CqSyncTimeColumn.Name)[0]
	//cqIDIndex := table.FieldIndices(schema.CqIDColumn.Name)[0]
	//sort.Slice(records, func(i, j int) bool {
	//	// sort by sync time, then UUID
	//	first := records[i].Column(syncTimeIndex).(*array.Timestamp).Value(0).ToTime(arrow.Millisecond)
	//	second := records[j].Column(syncTimeIndex).(*array.Timestamp).Value(0).ToTime(arrow.Millisecond)
	//	if first.Equal(second) {
	//		firstUUID := records[i].Column(cqIDIndex).(*types.UUIDArray).Value(0).String()
	//		secondUUID := records[j].Column(cqIDIndex).(*types.UUIDArray).Value(0).String()
	//		return strings.Compare(firstUUID, secondUUID) < 0
	//	}
	//	return first.Before(second)
	//})
	//
	//for i := range records {
	//	res <- records[i]
	//}
	return rs.GetError()
}
