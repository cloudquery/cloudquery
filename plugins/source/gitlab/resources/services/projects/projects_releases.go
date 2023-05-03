package projects

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func ProjectsReleases() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_projects_releases",
		Resolver:  fetchProjectsReleases,
		Transform: client.TransformWithStruct(&gitlab.Release{}, transformers.WithPrimaryKeys("CreatedAt")),
		Columns: schema.ColumnList{client.BaseURLColumn,
			{
				Name:            "project_id",
				Type:            schema.TypeInt,
				Resolver:        resolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{NotNull: true, PrimaryKey: true},
			},
		},
	}
}
