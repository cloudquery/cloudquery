package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
	"github.com/thoas/go-funk"
)

func fetchServices(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		p := c.Fastly.NewListServicesPaginator(&fastly.ListServicesInput{
			PerPage: 100,
		})
		cfg := c.Spec
		if p.HasNext() {
			services, err := p.GetNext()
			if err != nil {
				return err
			}
			for _, s := range services {
				if len(cfg.Services) > 0 && !funk.ContainsString(cfg.Services, s.ID) {
					continue
				}
				res <- s
			}
		}
		return nil
	}
	return c.RetryOnError(ctx, "fastly_services", f)
}
