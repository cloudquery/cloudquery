package engage

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EngageRevenues() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_engage_revenues",
		Resolver:  fetchEngageRevenues,
		Transform: transformers.TransformWithStruct(&mixpanel.EngageRevenue{}, client.SharedTransformers(transformers.WithPrimaryKeys("Date"))...),
		Columns: schema.ColumnList{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
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
