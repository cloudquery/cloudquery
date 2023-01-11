package worker_routes

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func WorkerRoutes() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_routes",
		Resolver:  fetchWorkerRoutes,
		Multiplex: client.ZoneMultiplex,
		Transform: transformers.TransformWithStruct(&cloudflare.WorkerRoute{}),
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
		},
	}
}
