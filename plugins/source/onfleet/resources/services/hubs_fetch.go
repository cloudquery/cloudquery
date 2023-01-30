package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchHubs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	result, err := cqClient.OnfleetClient.Hubs().List(ctx)
	if err != nil {
		return err
	}

	res <- result

	return nil
}
