package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamPrincipals(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	input := &ram.ListPrincipalsInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: types.ResourceOwnerSelf,
	}
	paginator := ram.NewListPrincipalsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Principals
	}

	input = &ram.ListPrincipalsInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: types.ResourceOwnerOtherAccounts,
	}
	paginator = ram.NewListPrincipalsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Principals
	}

	return nil
}
