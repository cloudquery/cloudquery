package registries

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchRegistriesRegistries(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	getFunc := func() error {
		response, _, err := svc.Services.Registry.Get(ctx)
		if err != nil {
			return err
		}
		res <- response
		return nil
	}

	err := client.ThrottleWrapper(ctx, svc, getFunc)
	if err != nil {
		return err
	}
	return nil
}
