package funnels

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Funnels() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_funnels",
		Resolver:  fetchFunnels,
		Transform: transformers.TransformWithStruct(&mixpanel.Funnel{}, client.SharedTransformers(transformers.WithPrimaryKeys("id"))...),
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
