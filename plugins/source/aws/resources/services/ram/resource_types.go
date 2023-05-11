package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func ResourceTypes() *schema.Table {
	tableName := "aws_ram_resource_types"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html`,
		Resolver:    fetchRamResourceTypes,
		Transform: transformers.TransformWithStruct(&types.ServiceNameAndResourceType{},
			transformers.WithPrimaryKeys("ResourceType", "ServiceName")),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

func fetchRamResourceTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	input := &ram.ListResourceTypesInput{MaxResults: aws.Int32(500)}
	paginator := ram.NewListResourceTypesPaginator(meta.(*client.Client).Services().Ram, input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *ram.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.ResourceTypes
	}
	return nil
}
