// Code generated by codegen; DO NOT EDIT.

package hooks

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Hooks() *schema.Table {
	return &schema.Table{
		Name:      "github_hooks",
		Resolver:  fetchHooks,
		Multiplex: client.OrgMultiplex,
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
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("URL"),
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "test_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TestURL"),
			},
			{
				Name:     "ping_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PingURL"),
			},
			{
				Name:     "last_response",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LastResponse"),
			},
			{
				Name:     "config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Config"),
			},
			{
				Name:     "events",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Events"),
			},
			{
				Name:     "active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Active"),
			},
		},

		Relations: []*schema.Table{
			Deliveries(),
		},
	}
}
