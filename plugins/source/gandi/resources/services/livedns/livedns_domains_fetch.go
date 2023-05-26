package livedns

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

func fetchLiveDNSDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	list, err := cl.Services.LiveDNSClient.ListDomains()
	if err != nil {
		return err
	}
	res <- list
	return nil
}
