package certificate_packs

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchCertificatePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	zoneId := svc.ZoneId

	packs, err := svc.ClientApi.ListCertificatePacks(ctx, zoneId)
	if err != nil {
		return err
	}
	res <- packs
	return nil
}
