package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResourceShareAssociations() *schema.Table {
	tableName := "aws_ram_resource_share_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceShareAssociation.html`,
		Resolver:    fetchRamResourceShareAssociations,
		Transform:   transformers.TransformWithStruct(&types.ResourceShareAssociation{}, transformers.WithPrimaryKeys("AssociatedEntity", "ResourceShareArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}

func fetchRamResourceShareAssociations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
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

func fetchRamResourceShareAssociationsByType(ctx context.Context, meta schema.ClientMeta, resourceShareInput *ram.GetResourceShareAssociationsInput, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := ram.NewGetResourceShareAssociationsPaginator(meta.(*client.Client).Services().Ram, resourceShareInput)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ResourceShareAssociations
	}
	return nil
}
