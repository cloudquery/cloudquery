package services

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func TestSomeTable() *schema.Table {
	return &schema.Table{
		Name:        "test_some_table",
		Description: "Test description",
		Resolver:    fetchSomeTableData,
		Columns: []schema.Column{
			{
				Name:        "column1",
				Description: "Test Column 1",
				Type:        arrow.BinaryTypes.String,
				PrimaryKey:  true,
			},
			{
				Name:        "column2",
				Description: "Test Column 2",
				Type:        arrow.PrimitiveTypes.Int64,
			},
		},
	}
}

func fetchSomeTableData(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	res <- map[string]any{
		"column1": "test_project_id",
		"column2": "test_id",
	}
	return nil
}
