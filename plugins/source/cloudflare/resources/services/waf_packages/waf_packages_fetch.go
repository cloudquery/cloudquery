package waf_packages

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchWAFPackages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFPackages(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWAFGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId
	pack := parent.Item.(cloudflare.WAFPackage)

	resp, err := svc.ClientApi.ListWAFGroups(ctx, zoneId, pack.ID)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func fetchWAFRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId
	pack := parent.Item.(cloudflare.WAFPackage)

	resp, err := svc.ClientApi.ListWAFRules(ctx, zoneId, pack.ID)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
