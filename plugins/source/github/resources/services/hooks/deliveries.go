// Code generated by codegen; DO NOT EDIT.

package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Deliveries() *schema.Table {
	return &schema.Table{
		Name:     "github_hook_deliveries",
		Resolver: fetchDeliveries,
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "hook_id",
				Type:        schema.TypeInt,
				Resolver:    client.ResolveParentColumn("ID"),
				Description: `Hook ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "delivered_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeliveredAt.Time"),
			},
			{
				Name:     "guid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GUID"),
			},
			{
				Name:     "redelivery",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Redelivery"),
			},
			{
				Name:     "duration",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("Duration"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "status_code",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StatusCode"),
			},
			{
				Name:     "event",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Event"),
			},
			{
				Name:     "action",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Action"),
			},
			{
				Name:     "installation_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InstallationID"),
			},
			{
				Name:     "repository_id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RepositoryID"),
			},
			{
				Name:     "request",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Request"),
			},
			{
				Name:     "response",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Response"),
			},
		},
	}
}
