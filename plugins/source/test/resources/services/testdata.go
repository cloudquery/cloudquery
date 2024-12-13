package services

import (
	"context"

	"github.com/apache/arrow-go/v18/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestDataTable() *schema.Table {
	table := schema.TestTable("test_testdata_table", schema.TestSourceOptions{
		SkipMaps: true,
	})

	idIndex := -1
	for i := range table.Columns {
		if table.Columns[i].Name == `id` {
			idIndex = i
		}
		table.Columns[i].PrimaryKey = false
		table.Columns[i].IncrementalKey = false
		table.Columns[i].NotNull = false
		table.Columns[i].Resolver = schema.PathResolver(table.Columns[i].Name)
	}
	if idIndex > -1 {
		table.Columns = append(table.Columns[:idIndex], table.Columns[idIndex+1:]...)
	}

	table.Columns = append(table.Columns, schema.Column{
		Name:        "client_id",
		Description: "ID of client",
		Type:        arrow.PrimitiveTypes.Int64,
		Resolver:    client.ResolveClientID,
	})

	tg := schema.NewTestDataGenerator(0)
	data := tg.Generate(table, schema.GenTestDataOptions{
		MaxRows: 1,
	})
	if data.NumRows() != 1 {
		panic("Expected 1 row of data")
	}

	dataAsMap := make(map[string]any)
	for i, col := range data.Columns() {
		dataAsMap[data.ColumnName(i)] = col.ValueStr(0)
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
