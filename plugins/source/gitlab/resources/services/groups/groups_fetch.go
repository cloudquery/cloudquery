package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	opt := &gitlab.ListGroupsOptions{
		MinAccessLevel: c.MinAccessLevel,
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
		// Get the first page with projects.
		groups, resp, err := c.Gitlab.Groups.ListGroups(opt, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
		if len(groups) == 0 {
			return nil
		}
		res <- groups

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

	return nil
}
