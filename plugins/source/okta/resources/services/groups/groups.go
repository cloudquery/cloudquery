package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "okta_groups",
		Resolver:  fetchGroups,
		Transform: client.TransformWithStruct(&okta.Group{}),
		Relations: schema.Tables{users()},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) (err error) {
	defer func() {
		err = client.ProcessOktaAPIError(err)
	}()

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
