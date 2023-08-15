package projects

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/xanzy/go-gitlab"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_projects",
		Resolver:  fetchProjects,
		Transform: client.TransformWithStruct(&gitlab.Project{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
		Relations: schema.Tables{releases(), branches(), members()},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	opt := &gitlab.ListProjectsOptions{
		MinAccessLevel: c.MinAccessLevel,
		ListOptions: gitlab.ListOptions{
			PerPage: 1000,
		},
	}
	for {
		projects, resp, err := c.Gitlab.Projects.ListProjects(opt, gitlab.WithContext(ctx))
		if err != nil {
			return err
		}
		res <- projects

		// Exit the loop when we've seen all pages.
		if resp.NextPage == 0 {
			break
		}

		// Update the page number to get the next page.
		opt.Page = resp.NextPage
	}

	return nil
}

var projectIDColumn = schema.Column{
	Name:       "project_id",
	Type:       arrow.PrimitiveTypes.Int64,
	NotNull:    true,
	PrimaryKey: true,
	Resolver: func(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		return resource.Set(c.Name, resource.Parent.Item.(*gitlab.Project).ID)
	},
}
