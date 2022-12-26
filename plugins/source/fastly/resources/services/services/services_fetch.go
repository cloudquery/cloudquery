package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServices(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	p := c.Fastly.NewListServicesPaginator(&fastly.ListServicesInput{})
	if p.HasNext() {
		services, err := p.GetNext()
		if err != nil {
			return err
		}
		res <- services
	}
	return nil
}

// Doesn't seem to provide any additional information?
//
//func getFastlyServiceDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
//	r := resource.Item.(fastly.Service)
//	svc, err := meta.(*client.Client).Fastly.GetServiceDetails(&fastly.GetServiceInput{
//		ID: r.ID,
//	})
//	if err != nil {
//		return err
//	}
//	resource.SetItem(svc)
//	return nil
//}
