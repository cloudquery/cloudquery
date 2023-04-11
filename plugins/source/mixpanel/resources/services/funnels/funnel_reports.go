package funnels

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func FunnelReports() *schema.Table {
	return &schema.Table{
		Name:        "mixpanel_funnel_reports",
		Description: `https://developer.mixpanel.com/reference/funnels-query`,
		Resolver:    fetchFunnelReports,
		Transform:   client.TransformWithStruct(&mixpanel.FunnelData{}, transformers.WithPrimaryKeys("Date")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
			},
			{
				Name:     "funnel_id",
				Type:     schema.TypeInt,
				Resolver: schema.ParentColumnResolver("funnel_id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchFunnelReports(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	f := parent.Item.(mixpanel.Funnel)

	ret, err := cl.Services.QueryFunnel(ctx, f.FunnelID, cl.MPSpec.StartDate, cl.MPSpec.EndDate)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}
