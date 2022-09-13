package services

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FetchWAFPackages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	resp, err := svc.ClientApi.ListWAFPackages(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- resp

	return nil
}
func FetchWAFGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func FetchWAFRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
