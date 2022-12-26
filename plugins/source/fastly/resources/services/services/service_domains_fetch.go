package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServiceDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	v := parent.Item.(*fastly.Version)
	domains, err := c.Fastly.ListDomains(&fastly.ListDomainsInput{
		ServiceID:      v.ServiceID,
		ServiceVersion: v.Number,
	})
	if err != nil {
		return err
	}
	res <- domains
	return nil
}
