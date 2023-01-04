package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ApiIntegrationResponses() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_api_integration_responses",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_IntegrationResponse.html`,
		Resolver:    fetchApigatewayv2ApiIntegrationResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.IntegrationResponse{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "api_integration_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "integration_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("integration_id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiIntegrationResponseArn(),
			},
		},
	}
}
