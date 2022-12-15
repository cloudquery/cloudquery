package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)

	req := cl.Services.Users.ListUsers(ctx).Limit(200)
	users, resp, err := cl.Services.Users.ListUsersExecute(req)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	res <- users

	for resp != nil && resp.HasNextPage() {
		var nextUserSet []okta.User
		resp, err = resp.Next(&nextUserSet)
		if err != nil {
			return err
		}
		res <- nextUserSet
	}
	return nil
}
