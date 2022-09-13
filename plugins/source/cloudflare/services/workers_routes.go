package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func WorkersRoutes() *schema.Table {
	return &schema.Table{
		Name:        "cloudflare_workers_routes",
		Description: "WorkerRoute is used to map traffic matching a URL pattern to a workers  API reference: https://api.cloudflare.com/#worker-routes-properties",
		Resolver:    fetchWorkersRoutes,
		Multiplex:   client.ZoneMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountID,
			},
			{
				Name:        "zone_id",
				Description: "The Zone ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneID,
			},
			{
				Name:            "id",
				Description:     "API item identifier tag",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "pattern",
				Description: "The pattern of the route.",
				Type:        schema.TypeString,
			},
			{
				Name:        "enabled",
				Description: "Whether the route is enabled",
				Type:        schema.TypeBool,
			},
			{
				Name:        "script",
				Description: "Name of the script to apply when the route is matched. The route is skipped when this is blank/missing.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchWorkersRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWorkerRoutes(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp.Routes

	return nil
}
