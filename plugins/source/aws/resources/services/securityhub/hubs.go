package securityhub

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Hubs() *schema.Table {
	tableName := "aws_securityhub_hubs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html`,
		Resolver:    fetchHubs,
		Transform: transformers.TransformWithStruct(&securityhub.DescribeHubOutput{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithPrimaryKeys("HubArn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: fetchHubTags,
			},
		},
	}
}

func fetchHubs(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Securityhub
	hub, err := svc.DescribeHub(ctx, nil, func(o *securityhub.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- hub
	return nil
}

func fetchHubTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Securityhub
	config := &securityhub.ListTagsForResourceInput{ResourceArn: resource.Item.(*securityhub.DescribeHubOutput).HubArn}
	tags, err := svc.ListTagsForResource(ctx, config, func(o *securityhub.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, tags.Tags)
}
