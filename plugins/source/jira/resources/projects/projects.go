package projects

import (
	"context"

	"github.com/andygrunwald/go-jira"
	"github.com/cloudquery/cloudquery/plugins/source/jira/sync"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type ProjectListEntry struct {
	Expand          string               `json:"expand" structs:"expand"`
	Self            string               `json:"self" structs:"self"`
	ID              string               `json:"id" structs:"id"`
	Key             string               `json:"key" structs:"key"`
	Name            string               `json:"name" structs:"name"`
	AvatarUrls      jira.AvatarUrls      `json:"avatarUrls" structs:"avatarUrls"`
	ProjectTypeKey  string               `json:"projectTypeKey" structs:"projectTypeKey"`
	ProjectCategory jira.ProjectCategory `json:"projectCategory,omitempty" structs:"projectsCategory,omitempty"`
	IssueTypes      []jira.IssueType     `json:"issueTypes,omitempty" structs:"issueTypes,omitempty"`
}

func Projects() *schema.Table {
	return &schema.Table{
		Name:      "jira_projects",
		Transform: transformers.TransformWithStruct(&ProjectListEntry{}, transformers.WithPrimaryKeys("Self")),
		Resolver:  fetchProjects,
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*sync.Client)
	projectList, _, err := c.Jira.Project.GetListWithContext(ctx)
	if err != nil {
		return err
	}
	res <- projectList
	return nil
}
