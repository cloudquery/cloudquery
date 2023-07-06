package groups

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/xanzy/go-gitlab"
)

func billableMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_group_billable_members",
		Resolver:  fetchBillableMembers,
		Transform: client.TransformWithStruct(new(gitlab.BillableGroupMember), transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn, groupIDColumn},
	}
}

func fetchBillableMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	group := parent.Item.(*gitlab.Group)
	opt := &gitlab.ListBillableGroupMembersOptions{ListOptions: gitlab.ListOptions{PerPage: 1000}}

	for {
		members, resp, err := c.Gitlab.Groups.ListBillableGroupMembers(group.ID, opt, gitlab.WithContext(ctx))
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
