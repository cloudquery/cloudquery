package ram

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ResourceTypes() *schema.Table {
	tableName := "aws_ram_resource_types"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html`,
		Resolver:    fetchRamResourceTypes,
		Transform: transformers.TransformWithStruct(&types.ServiceNameAndResourceType{},
			transformers.WithPrimaryKeyComponents("ResourceType", "ServiceName")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			{
				Name:                "account_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            client.ResolveAWSAccount,
				PrimaryKeyComponent: true,
			},
			{
				Name:                "region",
				Type:                arrow.BinaryTypes.String,
				Resolver:            client.ResolveAWSRegion,
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchRamResourceTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	input := &ram.ListResourceTypesInput{MaxResults: aws.Int32(500)}
	paginator := ram.NewListResourceTypesPaginator(meta.(*client.Client).Services(client.AWSServiceRam).Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.ResourceTypes
	}
	return nil
}
