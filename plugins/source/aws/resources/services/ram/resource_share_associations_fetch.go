package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamResourceShareAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	err := fetchRamResourceShareAssociationsByType(ctx, meta, &ram.GetResourceShareAssociationsInput{
		AssociationType: types.ResourceShareAssociationTypeResource,
		MaxResults:      aws.Int32(500),
	}, res)
	if err != nil {
		return err
	}

	err = fetchRamResourceShareAssociationsByType(ctx, meta, &ram.GetResourceShareAssociationsInput{
		AssociationType: types.ResourceShareAssociationTypePrincipal,
		MaxResults:      aws.Int32(500),
	}, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamResourceShareAssociationsByType(ctx context.Context, meta schema.ClientMeta, resourceShareInput *ram.GetResourceShareAssociationsInput, res chan<- interface{}) error {
	paginator := ram.NewGetResourceShareAssociationsPaginator(meta.(*client.Client).Services().Ram, resourceShareInput)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ResourceShareAssociations
	}
	return nil

}
