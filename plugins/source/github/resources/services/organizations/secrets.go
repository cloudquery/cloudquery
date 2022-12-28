// Code generated by codegen; DO NOT EDIT.

package organizations

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Secrets() *schema.Table {
	return &schema.Table{
		Name:     "github_organization_dependabot_secrets",
		Resolver: fetchSecrets,
		Columns: []schema.Column{
			{
				Name:        "org",
				Type:        schema.TypeString,
				Resolver:    client.ResolveOrg,
				Description: `The Github Organization of the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "visibility",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Visibility"),
			},
			{
				Name:     "selected_repositories_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelectedRepositoriesURL"),
			},
		},
	}
}
