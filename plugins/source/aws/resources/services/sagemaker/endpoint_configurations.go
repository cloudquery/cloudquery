package sagemaker

import (
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func EndpointConfigurations() *schema.Table {
	return &schema.Table{
		Name:                "aws_sagemaker_endpoint_configurations",
		Description:         `https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html`,
		Resolver:            fetchSagemakerEndpointConfigurations,
		PreResourceResolver: getEndpointConfiguration,
		Transform:           transformers.TransformWithStruct(&sagemaker.DescribeEndpointConfigOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("api.sagemaker"),
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
