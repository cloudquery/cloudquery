package services

import (
	"context"

	"github.com/cloudquery/cq-provider-cloudflare/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource workers_routes --config workers.hcl --output .
func WorkersRoutes() *schema.Table {
	return &schema.Table{
		Name:         "cloudflare_workers_routes",
		Description:  "WorkerRoute is used to map traffic matching a URL pattern to a workers  API reference: https://api.cloudflare.com/#worker-routes-properties",
		Resolver:     fetchWorkersRoutes,
		Multiplex:    client.ZoneMultiplex,
		DeleteFilter: client.DeleteAccountZoneFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountId,
			},
			{
				Name:        "zone_id",
				Description: "The Zone ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveZoneId,
			},
			{
				Name:        "id",
				Description: "API item identifier tag",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
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
		return diag.WrapError(err)
	}
	res <- resp.Routes

	return nil
}
