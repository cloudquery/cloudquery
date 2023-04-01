package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
