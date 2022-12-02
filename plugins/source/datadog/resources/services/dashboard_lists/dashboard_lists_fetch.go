package dashboard_lists

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDashboardLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	ctx = c.BuildContextV1(ctx)
	resp, _, err := c.DDServices.DashboardListsAPI.ListDashboardLists(ctx)
	if err != nil {
		return err
	}
	res <- resp.GetDashboardLists()
	return nil
}
