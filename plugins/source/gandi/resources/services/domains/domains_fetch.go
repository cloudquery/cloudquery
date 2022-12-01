package domains

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/go-gandi/go-gandi/domain"
)

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	list, err := cl.Services.DomainClient.ListDomains()
	if err != nil {
		return err
	}
	res <- list
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services.DomainClient
	output, err := svc.GetDomain(resource.Item.(domain.ListResponse).FQDN)
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
