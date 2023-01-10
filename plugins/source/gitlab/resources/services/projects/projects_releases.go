package projects

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/xanzy/go-gitlab"
)

func ProjectsReleases() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_projects_releases",
		Resolver:  fetchProjectsReleases,
		Transform: transformers.TransformWithStruct(&gitlab.Release{}, client.SharedTransformers()...),
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
				Name:     "project_id",
				Type:     schema.TypeInt,
				Resolver: resolveProjectID,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
