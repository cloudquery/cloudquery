package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func GroupMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_group_members",
		Resolver:  fetchGroupMembers,
		Transform: client.TransformWithStruct(&gitlab.GroupMember{}, transformers.WithPrimaryKeys("ID")),
		Columns: schema.ColumnList{client.BaseURLColumn,
			{
				Name:            "group_id",
				Type:            schema.TypeInt,
				Resolver:        resolveGroupID,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, PrimaryKey: true},
			},
		},
	}
}

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

func resolveGroupID(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	return resource.Set(c.Name, resource.Parent.Item.(*gitlab.Group).ID)
}
