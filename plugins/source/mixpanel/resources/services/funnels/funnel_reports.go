package funnels

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: client.ResolveProjectID,
			},
			{
				Name:       "funnel_id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.ParentColumnResolver("funnel_id"),
				PrimaryKey: true,
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
