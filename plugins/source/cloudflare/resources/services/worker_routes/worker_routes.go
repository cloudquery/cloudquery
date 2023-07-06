package worker_routes

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func WorkerRoutes() *schema.Table {
	return &schema.Table{
		Name:      "cloudflare_worker_routes",
		Resolver:  fetchWorkerRoutes,
		Multiplex: client.ZoneMultiplex,
		Transform: client.TransformWithStruct(&cloudflare.WorkerRoute{}),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveAccountID,
				Description: `The Account ID of the resource.`,
			},
			{
				Name:        "zone_id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    client.ResolveZoneID,
				Description: `Zone identifier tag.`,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
