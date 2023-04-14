package simplehosting

import (
	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func SimplehostingInstances() *schema.Table {
	return &schema.Table{
		Name:     "gandi_simplehosting_instances",
		Resolver: fetchSimplehostingInstances,
		Columns: []schema.Column{
			{
				Name:        "sharing_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveSharingID,
				Description: `The Sharing ID of the resource.`,
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "size",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Size"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "database",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Database"),
			},
			{
				Name:     "language",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Language"),
			},
			{
				Name:     "datacenter",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Datacenter"),
			},
		},

		Relations: []*schema.Table{
			SimplehostingInstanceVhosts(),
		},
	}
}
