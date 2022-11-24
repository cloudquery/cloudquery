package livedns

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/livedns"
)

func fetchLiveDNSDomainSnapshots(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	dom := parent.Item.(livedns.Domain)

	list, err := cl.Services.LiveDNSClient.ListSnapshots(dom.FQDN)
	if err != nil {
		return err
	}
	res <- list
	return nil
}
