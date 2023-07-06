package users

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/xanzy/go-gitlab"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&gitlab.User{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	if c.BaseURL == "" {
		c.Logger().Info().Str("table", "gitlab_users").Msg("not supported for GitLab SaaS, skipping...")
		return nil
	}

	opt := &gitlab.ListUsersOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}
	for {
		users, resp, err := c.Gitlab.Users.ListUsers(opt, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}

		res <- users

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}
	return nil
}
