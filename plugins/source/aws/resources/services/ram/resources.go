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

func Resources() *schema.Table {
	tableName := "aws_ram_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html`,
		Resolver:    fetchRamResources,
		Transform:   transformers.TransformWithStruct(&types.Resource{}, transformers.WithPrimaryKeys("Arn", "ResourceShareArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchRamResources(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	err := fetchRamResourcesByOwner(ctx, meta, types.ResourceOwnerSelf, res)
	if err != nil {
		return err
	}
	err = fetchRamResourcesByOwner(ctx, meta, types.ResourceOwnerOtherAccounts, res)
	if err != nil {
		return err
	}
	return nil
}

func fetchRamResourcesByOwner(ctx context.Context, meta schema.ClientMeta, shareType types.ResourceOwner, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &ram.ListResourcesInput{
		MaxResults:    aws.Int32(500),
		ResourceOwner: shareType,
	}
	paginator := ram.NewListResourcesPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Resources
	}
	return nil
}
