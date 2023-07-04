package resources

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:     "{{.Name}}_sample_table",
		Resolver: fetchSampleTable,
		Columns: []schema.Column{
      {
        Name: "column",
        Type: arrow.BinaryTypes.String,
      },
    },
  }
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan <- any) error {
  return fmt.Errorf("not implemented")
}
