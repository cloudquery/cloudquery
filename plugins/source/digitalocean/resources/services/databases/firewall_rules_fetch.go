package databases

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/digitalocean/godo"
)

func fetchDatabasesFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(godo.Database)
	svc := meta.(*client.Client)
	getFunc := func() error {
		response, _, err := svc.Services.Databases.GetFirewallRules(ctx, p.ID)
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
