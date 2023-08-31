package waf_overrides

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchWAFOverrides(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFOverrides(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
