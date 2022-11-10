package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRamResourceShareAssociatedResources(
	ctx context.Context,
	meta schema.ClientMeta,
	parent *schema.Resource,
	res chan<- interface{},
) error {
	input := &ram.GetResourceShareAssociationsInput{
		AssociationType: types.ResourceShareAssociationTypeResource,
		ResourceArn:     parent.Item.(types.ResourceShare).ResourceShareArn,
		MaxResults:      aws.Int32(500),
	}

	paginator := ram.NewGetResourceShareAssociationsPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.ResourceShareAssociations
	}
	return nil
}
