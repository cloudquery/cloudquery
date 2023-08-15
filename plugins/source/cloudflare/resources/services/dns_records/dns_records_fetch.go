package dns_records

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchDNSRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	records, err := svc.ClientApi.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{})
	if err != nil {
		return err
	}
	res <- records
	return nil
}
