// Code generated by codegen; DO NOT EDIT.

package deployment

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeploymentChecks() *schema.Table {
	return &schema.Table{
		Name:     "vercel_deployment_checks",
		Resolver: fetchDeploymentChecks,
		Columns: []schema.Column{
			{
				Name:     "deployment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("uid"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
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
				Name:     "completed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CompletedAt"),
			},
			{
				Name:     "conclusion",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Conclusion"),
			},
			{
				Name:     "details_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DetailsURL"),
			},
			{
				Name:     "integration_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IntegrationID"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
			{
				Name:     "rererequestable",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Rerequestable"),
			},
			{
				Name:     "started_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartedAt"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
