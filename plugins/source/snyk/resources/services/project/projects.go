package project

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Projects() *schema.Table {
	return &schema.Table{
		Name:        "snyk_projects",
		Description: `https://pkg.go.dev/github.com/pavel-snyk/snyk-sdk-go/snyk#Project`,
		Resolver:    fetchProjects,
		Multiplex:   client.ByOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Project{}),
		Columns: []schema.Column{
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveOrganizationID,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
