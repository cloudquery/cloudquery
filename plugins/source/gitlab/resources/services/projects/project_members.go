package projects

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func ProjectMembers() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_project_members",
		Resolver:  fetchProjectMembers,
		Transform: client.TransformWithStruct(&gitlab.ProjectMember{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.BaseURLColumn,
			{
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: resolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
					NotNull:    true,
				},
			},
		},
	}
}

func fetchProjectMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	project := parent.Item.(*gitlab.Project)
	opt := &gitlab.ListProjectMembersOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}

	for {
		// Get the first page with projects.
		members, resp, err := c.Gitlab.ProjectMembers.ListProjectMembers(project.ID, opt, gitlab.WithContext(ctx))
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
