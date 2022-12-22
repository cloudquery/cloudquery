package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func fetchGroupMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	group := parent.Item.(*gitlab.Group)
	opt := &gitlab.ListGroupMembersOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
		// Get the first page with projects.
		members, resp, err := c.Gitlab.Groups.ListGroupMembers(group.ID, opt, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
		res <- members
		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

	return nil
}

func resolveGroupID(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set("group_id", resource.Parent.Item.(*gitlab.Group).ID)
}
