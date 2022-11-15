package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamResourceShares(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	err := fetchRamResourceSharesByType(ctx, meta, types.ResourceOwnerOtherAccounts, res)
	if err != nil {
		return err
	}
	err = fetchRamResourceSharesByType(ctx, meta, types.ResourceOwnerSelf, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamResourceSharesByType(ctx context.Context, meta schema.ClientMeta, shareType types.ResourceOwner, res chan<- interface{}) error {
	input := &ram.GetResourceSharesInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: shareType,
	}
	paginator := ram.NewGetResourceSharesPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ResourceShares
	}
	return nil
}
