package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchAdaccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	item, err := cqClient.FacebookClient.GetAdaccount(ctx)

	if err != nil {
		return err
	}

	res <- item

	return nil
}
