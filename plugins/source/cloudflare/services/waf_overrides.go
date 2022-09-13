package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FetchWAFOverrides(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFOverrides(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
