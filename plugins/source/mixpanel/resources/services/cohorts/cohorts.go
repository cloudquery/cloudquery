package cohorts

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Cohorts() *schema.Table {
	return &schema.Table{
		Name:        "mixpanel_cohorts",
		Description: `https://developer.mixpanel.com/reference/cohorts-list`,
		Resolver:    fetchCohorts,
		Transform:   client.TransformWithStruct(&mixpanel.Cohort{}, transformers.WithPrimaryKeys("ID")),
		Relations: []*schema.Table{
			CohortMembers(),
		},
	}
}

func fetchCohorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	ret, err := cl.Services.ListCohorts(ctx)
	if err != nil {
		return err
	}
	res <- ret
	return nil
}
