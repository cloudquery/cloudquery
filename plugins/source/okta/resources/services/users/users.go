package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "okta_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&okta.User{}),
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) (err error) {
	defer func() {
		err = client.ProcessOktaAPIError(err)
	}()

	cl := meta.(*client.Client)

	req := cl.UserApi.ListUsers(ctx).Limit(200)
	users, resp, err := cl.UserApi.ListUsersExecute(req)
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
