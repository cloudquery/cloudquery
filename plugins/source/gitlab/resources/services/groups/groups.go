package groups

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/xanzy/go-gitlab"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_groups",
		Resolver:  fetchGroups,
		Transform: client.TransformWithStruct(&gitlab.Group{}, transformers.WithPrimaryKeys("ID", "Name")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
		Relations: schema.Tables{billableMembers(), members()},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	opt := &gitlab.ListGroupsOptions{
		MinAccessLevel: c.MinAccessLevel,
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
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

var groupIDColumn = schema.Column{
	Name:       "group_id",
	Type:       arrow.PrimitiveTypes.Int64,
	NotNull:    true,
	PrimaryKey: true,
	Resolver: func(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		return resource.Set(c.Name, resource.Parent.Item.(*gitlab.Group).ID)
	},
}
