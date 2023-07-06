package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/xanzy/go-gitlab"
)

func branches() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_project_branches",
		Resolver:  fetchBranches,
		Transform: client.TransformWithStruct(&gitlab.Branch{}, transformers.WithPrimaryKeys("Name")),
		Columns:   schema.ColumnList{client.BaseURLColumn, projectIDColumn},
	}
}

func fetchBranches(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	project := parent.Item.(*gitlab.Project)
	opt := &gitlab.ListBranchesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
		branches, resp, err := c.Gitlab.Branches.ListBranches(project.ID, opt, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
		res <- branches
		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

	return nil
}
