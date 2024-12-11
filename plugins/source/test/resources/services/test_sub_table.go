package services

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"golang.org/x/exp/maps"
)

func testSubTable(config client.Spec) *schema.Table {
	extraCols := make([]schema.Column, *config.NumSubCols)
	for i := 0; i < *config.NumSubCols; i++ {
		extraCols[i] = schema.Column{
			Name:        fmt.Sprintf("extra_column_%d", i),
			Description: fmt.Sprintf("Extra column %d", i),
			Type:        arrow.PrimitiveTypes.Int64,
		}
	}

	return &schema.Table{
		Name:        "test_sub_table",
		Description: "Sub table of test_some_table",
		Resolver:    fetchSubTableData,
		Columns: append([]schema.Column{
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
		}, extraCols...),
	}
}

func fetchSubTableData(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	if cl.Spec.FailImmediately {
		return ErrFailImmediately
	}
	if cl.Spec.ExitImmediately {
		os.Exit(1)
	}
	colMap := make(map[string]any, *cl.Spec.NumSubCols)
	for i := 0; i < *cl.Spec.NumSubCols; i++ {
		colMap["extra_column_"+strconv.FormatInt(int64(i), 10)] = rand.Int63()
	}

	for i := 0; i < *cl.Spec.NumSubRows; i++ {
		data := map[string]any{
			"sub_resource_id": i,
			"data_column":     fmt.Sprintf("sub_data_%d", i%3),
		}
		maps.Copy(data, colMap)
		res <- data
	}
	return nil
}
