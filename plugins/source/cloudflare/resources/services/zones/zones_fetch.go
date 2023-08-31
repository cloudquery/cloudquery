package zones

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)

	opts := cloudflare.WithZoneFilters("", svc.AccountId, "")

	resp, err := svc.ClientApi.ListZonesContext(ctx, opts)
	if err != nil {
		return err
	}
	res <- resp.Result

	return nil
}
