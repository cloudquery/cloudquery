package apps

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/digitalocean/godo"
)

func fetchAppsAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)
	app := parent.Item.(*godo.App)

	return client.ThrottleWrapper(ctx, svc, func() error {
		alerts, _, err := svc.Services.Apps.ListAlerts(ctx, app.ID)
		if err != nil {
			return err
		}
		res <- alerts
		return nil
	})
}
