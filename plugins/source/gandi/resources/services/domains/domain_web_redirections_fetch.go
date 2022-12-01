package domains

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/domain"
)

func fetchDomainWebRedirections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	dom := parent.Item.(domain.Details)

	list, err := cl.Services.DomainClient.ListWebRedirections(dom.FQDN)
	if err != nil {
		return err
	}
	res <- list
	return nil
}
