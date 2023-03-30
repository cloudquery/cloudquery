package projects

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_projects",
		Resolver:  fetchProjects,
		Transform: client.TransformWithStruct(&gitlab.Project{}),
		Columns: []schema.Column{
			{
				Name:     "base_url",
				Type:     schema.TypeString,
				Resolver: client.ResolveURL,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "marked_for_deletion_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("MarkedForDeletionAt"),
			},
		},

		Relations: []*schema.Table{
			ProjectsReleases(),
			ProjectBranches(),
		},
	}
}
