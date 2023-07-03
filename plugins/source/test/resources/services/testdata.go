package services

import (
	"context"
	"strings"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestDataTable() *schema.Table {
	table := schema.TestTable("test_testdata_table", schema.TestSourceOptions{
		SkipMaps: true,
	})
	for i, c := range table.Columns {
		if strings.HasPrefix(c.Name, "_cq_") {
			table.Columns[i].Name = "test" + c.Name
		}
		table.Columns[i].PrimaryKey = false
		table.Columns[i].IncrementalKey = false
		table.Columns[i].NotNull = false
		table.Columns[i].Resolver = schema.PathResolver(table.Columns[i].Name)
	}

	table.Columns = append(table.Columns, schema.Column{
		Name:        "client_id",
		Description: "ID of client",
		Type:        arrow.PrimitiveTypes.Int64,
		Resolver:    client.ResolveClientID,
	})

	data := schema.GenTestData(table, schema.GenTestDataOptions{
		MaxRows: 1,
	})
	if len(data) != 1 {
		panic("Expected 1 row of data")
	}

	dataAsMap := make(map[string]any)
	for i, col := range data[0].Columns() {
		dataAsMap[data[0].ColumnName(i)] = col.ValueStr(0)
	}

	table.Description = "Testdata table"
	table.Resolver = fetchTestData(dataAsMap)
	table.Multiplex = client.MultiplexBySpec
	return table
}

func fetchTestData(data map[string]any) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		res <- data
		return nil
	}
}
