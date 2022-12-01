package domains

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/domain"
)

func fetchDomainLiveDNS(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	dom := parent.Item.(domain.Details)

	output, err := cl.Services.DomainClient.GetLiveDNS(dom.FQDN)
	if err != nil {
		return err
	}
	res <- output
	return nil
}
