package projects

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/xanzy/go-gitlab"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_projects",
		Resolver:  fetchProjects,
		Transform: client.TransformWithStruct(&gitlab.Project{}, transformers.WithPrimaryKeys("ID")),
		Columns:   schema.ColumnList{client.BaseURLColumn},
		Relations: schema.Tables{ProjectsReleases(), ProjectBranches(), ProjectMembers()},
	}
}
