package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EndpointConfigurations() *schema.Table {
	tableName := "aws_sagemaker_endpoint_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html`,
		Resolver:            fetchSagemakerEndpointConfigurations,
		PreResourceResolver: getEndpointConfiguration,
		Transform:           transformers.TransformWithStruct(&sagemaker.DescribeEndpointConfigOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "api.sagemaker"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EndpointConfigArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveSagemakerEndpointConfigurationTags,
				Description: `The tags associated with the model.`,
			},
		},
	}
}

func fetchSagemakerEndpointConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListEndpointConfigsInput{}
	paginator := sagemaker.NewListEndpointConfigsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.EndpointConfigs
	}
	return nil
}

func getEndpointConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	n := resource.Item.(types.EndpointConfigSummary)

	response, err := svc.DescribeEndpointConfig(ctx, &sagemaker.DescribeEndpointConfigInput{
		EndpointConfigName: n.EndpointConfigName,
	})
	if err != nil {
		return err
	}

	resource.Item = response
	return nil
}

func resolveSagemakerEndpointConfigurationTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeEndpointConfigOutput)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.EndpointConfigArn,
	}
	paginator := sagemaker.NewListTagsPaginator(svc, &config)
	var tags []types.Tag
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set(col.Name, client.TagsToMap(tags))
}
