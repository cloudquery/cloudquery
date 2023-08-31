package project

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "snyk_projects",
		Description: `https://snyk.docs.apiary.io/#reference/projects/all-projects/list-all-projects`,
		Resolver:    fetchProjects,
		Multiplex:   client.ByOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Project{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.OrganizationID},
	}
}

func fetchProjects(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	projects, _, err := c.Projects.List(ctx, c.OrganizationID)
	if err != nil {
		return err
	}

	res <- projects

	return nil
}
