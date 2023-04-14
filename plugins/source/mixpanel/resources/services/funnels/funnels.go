package funnels

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Funnels() *schema.Table {
	return &schema.Table{
		Name:        "mixpanel_funnels",
		Description: `https://developer.mixpanel.com/reference/funnels-list-saved`,
		Resolver:    fetchFunnels,
		Transform:   client.TransformWithStruct(&mixpanel.Funnel{}, transformers.WithPrimaryKeys("FunnelID")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
			},
		},
		Relations: []*schema.Table{
			FunnelReports(),
		},
	}
}

func fetchFunnels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ListFunnels(ctx)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}
