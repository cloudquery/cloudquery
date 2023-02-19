package services

import (
	"context"
	"strings"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/testdata"
)

func TestDataTable() *schema.Table {
	table := testdata.TestSourceTable("test_testdata_table")
	for i, c := range table.Columns {
		if strings.HasPrefix(c.Name, "_cq_") {
			table.Columns[i].Name = "test" + c.Name
		}
		table.Columns[i].CreationOptions = schema.ColumnCreationOptions{}
		table.Columns[i].Resolver = schema.PathResolver(table.Columns[i].Name)
	}

	data := testdata.GenTestData(table)
	dataAsMap := make(map[string]any)
	for i, c := range table.Columns {
		if data[i].GetStatus() == schema.Present {
			dataAsMap[c.Name] = data[i].String()
		}
	}

	table.Description = "Testdata table"
	table.Resolver = fetchTestData(dataAsMap)
	return table
}

func fetchTestData(data map[string]any) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		res <- data
		return nil
	}
}
