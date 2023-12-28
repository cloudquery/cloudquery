package images

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/cloudquery/plugins/source/cloudflare/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	rc := cloudflare.AccountIdentifier(svc.AccountId)

	records, err := svc.ClientApi.ListImages(ctx, rc, cloudflare.ListImagesParams{})
	if err != nil {
		return err
	}
	res <- records
	return nil
}
