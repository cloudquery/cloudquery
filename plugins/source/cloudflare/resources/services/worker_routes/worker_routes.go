// Code generated by codegen; DO NOT EDIT.

package worker_routes

import (
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkerRoutes() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_routes",
		Resolver:  fetchWorkerRoutes,
		Multiplex: client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "zone_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneID,
				Description: `Zone identifier tag.`,
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
				Name:     "pattern",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Pattern"),
			},
			{
				Name:     "script",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScriptName"),
			},
		},
	}
}
