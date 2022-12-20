package applications

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func fetchApplicationUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	app := parent.Item.(*okta.Application)

	req := cl.ApplicationApi.ListApplicationUsers(ctx, *app.Id).Limit(200)
	items, resp, err := cl.ApplicationApi.ListApplicationUsersExecute(req)
	if err != nil {
		return err
	}
	if len(items) == 0 {
		return nil
	}
	res <- items

	for resp != nil && resp.HasNextPage() {
		var nextItems []*okta.AppUser
		resp, err = resp.Next(&nextItems)
		if err != nil {
			return err
		}
		res <- nextItems
	}
	return nil
}
