package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	req := cl.GroupApi.ListGroups(ctx).Limit(200)
	items, resp, err := cl.GroupApi.ListGroupsExecute(req)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items

	for resp != nil && resp.HasNextPage() {
		var nextItems []okta.Group
		resp, err = resp.Next(&nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
