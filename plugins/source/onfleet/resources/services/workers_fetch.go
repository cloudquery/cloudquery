package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/keplr-team/go-onfleet/onfleet"
)

func fetchWorkers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	workers, err := cqClient.OnfleetClient.Workers().List(ctx, &onfleet.WorkersListOptions{})
	if err != nil {
		return err
	}

	res <- workers

	return nil
}
