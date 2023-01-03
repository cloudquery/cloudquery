package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/xanzy/go-gitlab"
)

func Projects() []*Resource {
	resources := []*Resource{
		{
			Service:    "projects",
			SubService: "projects",
			PKColumns:  []string{"base_url", "id"},
			Struct:     &gitlab.Project{},
			Relations:  []string{"ProjectsReleases()", "ProjectBranches()"},
		},
		{
			Service:    "projects",
			SubService: "project_branches",
			Struct:     &gitlab.Branch{},
			PKColumns:  []string{"project_id", "name"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "project_id",
					Type:     schema.TypeInt,
					Resolver: `resolveProjectID`,
				},
			},
		},
		{
			Service:    "projects",
			SubService: "projects_releases",
			Struct:     &gitlab.Release{},
			PKColumns:  []string{"base_url", "project_id", "created_at"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "project_id",
					Type:     schema.TypeInt,
					Resolver: `resolveProjectID`,
				},
			},
		},
	}

	return resources
}
