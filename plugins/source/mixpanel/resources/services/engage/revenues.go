package engage

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func EngageRevenues() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_engage_revenues",
		Resolver:  fetchEngageRevenues,
		Transform: client.TransformWithStruct(&mixpanel.EngageRevenue{}, transformers.WithPrimaryKeys("Date")),
		Columns: schema.ColumnList{
			{
				Name:       "project_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   client.ResolveProjectID,
				PrimaryKey: true,
			},
		},
	}
}

func fetchEngageRevenues(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ListEngageRevenues(ctx, cl.MPSpec.StartDate, cl.MPSpec.EndDate)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}
