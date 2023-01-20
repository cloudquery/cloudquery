package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func TestSomeTable() *schema.Table {
	return &schema.Table{
		Name:        "test_some_table",
		Description: "Test description",
		Resolver:    fetchTestData,
		Columns: []schema.Column{
			{
				Name:            "column1",
				Description:     "Test Column 1",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "column2",
				Description: "Test Column 2",
				Type:        schema.TypeInt,
			},
		},
	}
}

func fetchTestData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	res <- map[string]any{
		"column1": "test_project_id",
		"column2": "test_id",
	}
	return nil
}
