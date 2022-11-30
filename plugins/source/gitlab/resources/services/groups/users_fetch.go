package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	groupMember := parent.Item.(*gitlab.GroupMember)

	opt := gitlab.GetUsersOptions{}

	// Get the first page with projects.
	users, _, err := c.Gitlab.Users.GetUser(groupMember.ID, opt)
	if err != nil {
		return err
	}

	res <- users

	return nil
}
