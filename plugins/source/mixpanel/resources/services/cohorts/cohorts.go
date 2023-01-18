package cohorts

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Cohorts() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_cohorts",
		Resolver:  fetchCohorts,
		Transform: transformers.TransformWithStruct(&mixpanel.Cohort{}, transformers.WithPrimaryKeys("id")),
	}
}

func fetchCohorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ListFunnels(ctx)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}
