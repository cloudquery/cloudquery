package cohorts

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CohortMembers() *schema.Table {
	return &schema.Table{
		Name:        "mixpanel_cohort_members",
		Description: `https://developer.mixpanel.com/reference/engage-query`,
		Resolver:    fetchCohortMembers,
		Transform:   transformers.TransformWithStruct(&mixpanel.EngageProfile{}, client.SharedTransformers(transformers.WithPrimaryKeys("DistinctID"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: client.ResolveProjectID,
			},
			{
				Name:     "cohort_id",
				Type:     schema.TypeInt,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCohortMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	co := parent.Item.(mixpanel.Cohort)
	qp := url.Values{}
	qp.Set("filter_by_cohort", fmt.Sprintf(`{"id":%d}`, co.ID))
	qp.Set("include_all_users", "true")

	var pg *mixpanel.EngagePaginator

	for {
		ret, err := cl.Services.EngageProfiles(ctx, qp)
		if err != nil {
			return err
		}
		res <- ret.Data

		if pg == nil {
			pg = &ret.EngagePaginator // Only page 0 has the paginator
		}
		page := ret.Page + 1
		if pg.TotalPages == 0 || page >= pg.TotalPages { // Zero based page numbers
			break
		}

		qp.Set("page", fmt.Sprintf("%d", page))
		qp.Set("session_id", ret.SessionID)
	}

	return nil
}
