package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	users, _, err := c.Gitlab.Users.ListUsers(&gitlab.ListUsersOptions{})
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return nil
	}
	res <- users
	return nil
}
