package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsmInventorySchemas(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm

	paginator := ssm.NewGetInventorySchemaPaginator(svc, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- v.Schemas
	}
	return nil
}
