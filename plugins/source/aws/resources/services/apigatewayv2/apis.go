package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Apis() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_apis",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Api.html`,
		Resolver:    fetchApigatewayv2Apis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:  transformers.TransformWithStruct(&types.Api{}),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiArn(),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
		},
		Relations: []*schema.Table{
			ApiAuthorizers(),
			ApiDeployments(),
			ApiIntegrations(),
			ApiModels(),
			ApiRoutes(),
			ApiStages(),
		},
	}
}
