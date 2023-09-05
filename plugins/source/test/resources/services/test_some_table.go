package services

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestSomeTable() *schema.Table {
	return &schema.Table{
		Name:        "test_some_table",
		Description: "Test table",
		Resolver:    fetchSomeTableData,
		Multiplex:   client.MultiplexBySpec,
		Columns: []schema.Column{
			{
				Name:        "resource_id",
				Description: "Resource ID",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    schema.PathResolver("resource_id"),
				PrimaryKey:  true,
			},
			{
				Name:        "column2",
				Description: "Test Column 2",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("column2"),
			},
			{
				Name:        "client_id",
				Description: "ID of client",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    client.ResolveClientID,
				PrimaryKey:  true,
			},
		},
		Relations: schema.Tables{
			testSubTable(),
		},
	}
}

func fetchSomeTableData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	for i := 0; i < *cl.Spec.NumRows; i++ {
		res <- map[string]any{
			"resource_id": i + 1,
			"column2":     fmt.Sprintf("test_column_data_%d", i%3),
		}
	}
	return nil
}
