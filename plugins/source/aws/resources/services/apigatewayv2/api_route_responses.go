package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func apiRouteResponses() *schema.Table {
	tableName := "aws_apigatewayv2_api_route_responses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RouteResponse.html`,
		Resolver:    fetchApigatewayv2ApiRouteResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.RouteResponse{}),
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

func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Route)
	p := parent.Parent.Item.(types.Api)
	config := apigatewayv2.GetRouteResponsesInput{
		ApiId:   p.ApiId,
		RouteId: r.RouteId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetRouteResponses(ctx, &config)

		if err != nil {
			return err
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveApiRouteResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.RouteResponse)
	route := resource.Parent.Item.(types.Route)
	p := resource.Parent.Parent.Item.(types.Api)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/apis/%s/routes/%s/routeresponses/%s", aws.ToString(p.ApiId), aws.ToString(route.RouteId), aws.ToString(r.RouteResponseId)),
	}.String())
}
