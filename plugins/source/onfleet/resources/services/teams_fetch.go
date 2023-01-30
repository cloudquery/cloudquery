package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	result, err := cqClient.OnfleetClient.Teams().List(ctx)
	if err != nil {
		return err
	}

	res <- result

	return nil
}
