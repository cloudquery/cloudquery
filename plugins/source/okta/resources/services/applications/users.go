package applications

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func users() *schema.Table {
	return &schema.Table{
		Name:      "okta_application_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&okta.AppUser{}),
		Columns:   schema.ColumnList{appIDColumn},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) (err error) {
	defer func() {
		err = client.ProcessOktaAPIError(err)
	}()

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
