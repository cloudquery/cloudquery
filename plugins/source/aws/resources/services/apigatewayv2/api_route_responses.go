package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiRouteResponses() *schema.Table {
	tableName := "aws_apigatewayv2_api_route_responses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html`,
		Resolver:    fetchApigatewayv2ApiRouteResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&types.RouteResponse{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "api_route_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "route_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("route_id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiRouteResponseArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
