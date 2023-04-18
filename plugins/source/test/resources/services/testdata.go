package services

import (
	"context"
	"strings"
	"time"

	"github.com/apache/arrow/go/v12/arrow/memory"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/testdata"
	"github.com/google/uuid"
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

	records := testdata.GenTestData(memory.DefaultAllocator, schema.CQSchemaToArrow(table), testdata.GenTestDataOptions{
		SourceName: "test",
		SyncTime:   time.Now(),
		MaxRows:    1,
		StableUUID: uuid.Nil,
	})
	defer func() {
		for _, record := range records {
			record.Release()
		}
	}()

	dataAsMap := make(map[string]any)
	for i, c := range table.Columns {
		dataAsMap[c.Name] = records[0].Column(i).String()
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
