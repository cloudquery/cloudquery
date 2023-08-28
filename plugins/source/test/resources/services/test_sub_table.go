package services

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func testSubTable() *schema.Table {
	return &schema.Table{
		Name:        "test_sub_table",
		Description: "Sub table of test_some_table",
		Resolver:    fetchSubTableData,
		Columns: []schema.Column{
			{
				Name:        "parent_resource_id",
				Description: "Parent resource ID",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    schema.ParentColumnResolver("resource_id"),
				PrimaryKey:  true,
			},
			{
				Name:        "sub_resource_id",
				Description: "Sub resource ID",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    schema.PathResolver("sub_resource_id"),
				PrimaryKey:  true,
			},
			{
				Name:        "data_column",
				Description: "Data column",
				Type:        arrow.BinaryTypes.String,
				Resolver:    schema.PathResolver("data_column"),
			},
		},
	}
}

func fetchSubTableData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	for i := 0; i < *cl.Spec.NumSubRows; i++ {
		res <- map[string]any{
			"sub_resource_id": i,
			"data_column":     fmt.Sprintf("sub_data_%d", i%3),
		}
	}
	return nil
}
