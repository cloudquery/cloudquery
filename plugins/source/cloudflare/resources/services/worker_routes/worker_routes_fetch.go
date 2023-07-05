package worker_routes

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchWorkerRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	rc := cloudflare.ZoneIdentifier(zoneId)
	params := cloudflare.ListWorkerRoutesParams{}
	resp, err := svc.ClientApi.ListWorkerRoutes(ctx, rc, params)
	if err != nil {
		return err
	}
	res <- resp.Routes

	return nil
}
