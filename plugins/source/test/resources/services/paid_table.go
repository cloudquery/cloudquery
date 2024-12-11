package services

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var ErrFailImmediately = errors.New("failing immediately")

func TestPaidTable() *schema.Table {
	return &schema.Table{
		Name:        "test_paid_table",
		Description: "Test Paid table",
		Resolver:    fetchPaidTableData,
		Multiplex:   client.MultiplexBySpec,
		IsPaid:      true,
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
	}
}

func fetchPaidTableData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	if cl.Spec.FailImmediately {
		return ErrFailImmediately
	}
	if cl.Spec.ExitImmediately {
		os.Exit(1)
	}
	for i := 0; i < *cl.Spec.NumRows; i++ {
		res <- map[string]any{
			"resource_id": i + 1,
			"column2":     fmt.Sprintf("test_column_data_%d", i%3),
		}
	}
	return nil
}
