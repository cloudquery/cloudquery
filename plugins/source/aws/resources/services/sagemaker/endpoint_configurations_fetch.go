package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSagemakerEndpointConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListEndpointConfigsInput{}
	for {
		response, err := svc.ListEndpointConfigs(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.EndpointConfigs

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
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

func resolveSagemakerEndpointConfigurationTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(*sagemaker.DescribeEndpointConfigOutput)
	c := meta.(*client.Client)
	svc := c.Services().Sagemaker
	config := sagemaker.ListTagsInput{
		ResourceArn: r.EndpointConfigArn,
	}
	response, err := svc.ListTags(ctx, &config)
	if err != nil {
		return err
	}

	tags := make(map[string]*string, len(response.Tags))
	for _, t := range response.Tags {
		tags[*t.Key] = t.Value
	}

	return resource.Set("tags", tags)
}
