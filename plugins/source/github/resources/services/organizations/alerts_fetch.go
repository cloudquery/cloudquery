package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAlerts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	alerts, _, err := c.Github.Dependabot.ListOrgAlerts(ctx, c.Org, nil)
	if err != nil {
		return err
	}

	res <- alerts

	return nil
}
