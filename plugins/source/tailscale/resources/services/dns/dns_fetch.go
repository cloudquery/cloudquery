package dns

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchDns(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, err := c.DNSPreferences(ctx)
	if err != nil {
		return err
	}

	res <- result
	return nil
}

func fetchDNSNameservers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)

	result, err := c.DNSNameservers(ctx)
	if err != nil {
		return err
	}

	return resource.Set(column.Name, result)
}

func fetchDNSSearchPaths(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, column schema.Column) error {
	c := meta.(*client.Client)

	result, err := c.DNSSearchPaths(ctx)
	if err != nil {
		return err
	}

	return resource.Set(column.Name, result)
}
