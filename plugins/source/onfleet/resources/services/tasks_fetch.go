package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/keplr-team/go-onfleet/onfleet"
)

func fetchTasks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	result, err := cqClient.OnfleetClient.Tasks().List(ctx, &onfleet.TasksListOptions{})
	if err != nil {
		return err
	}

	res <- result

	return nil
}
