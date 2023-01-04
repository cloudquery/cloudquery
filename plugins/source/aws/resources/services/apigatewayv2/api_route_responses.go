package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ApiRouteResponses() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigatewayv2_api_route_responses",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html`,
		Resolver:    fetchApigatewayv2ApiRouteResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:  transformers.TransformWithStruct(&types.RouteResponse{}),
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
				Name:     "api_route_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiRouteResponseArn(),
			},
		},
	}
}
