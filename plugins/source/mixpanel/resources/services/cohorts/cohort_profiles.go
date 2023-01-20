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

func CohortProfiles() *schema.Table {
	return &schema.Table{
		Name:      "mixpanel_cohort_profiles",
		Resolver:  fetchCohortProfiles,
		Transform: transformers.TransformWithStruct(&mixpanel.EngageProfile{}),
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

func fetchCohortProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	co := parent.Item.(mixpanel.Cohort)
	qp := url.Values{}
	qp.Add("filter_by_cohort", fmt.Sprintf(`{"id":%d}`, co.ID))

	page := int64(0)

	for {
		ret, err := cl.Services.EngageProfiles(ctx, qp)
		if err != nil {
			return err
		}
		res <- ret.Data

		page = ret.Page + 1
		if page > ret.TotalPages {
			break
		}

		qp.Set("page", fmt.Sprintf("%d", page))
		qp.Set("session_id", ret.SessionID)
	}

	return nil
}
