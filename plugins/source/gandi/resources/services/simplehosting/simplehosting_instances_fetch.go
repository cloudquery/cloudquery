package simplehosting

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/gandi/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
)

func fetchSimplehostingInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	list, err := cl.Services.SimpleHostingClient.ListInstances()
	if err != nil {
		return err
	}
	res <- list
	return nil
}
