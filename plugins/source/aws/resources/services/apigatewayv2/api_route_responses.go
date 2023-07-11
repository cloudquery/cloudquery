package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func apiRouteResponses() *schema.Table {
	tableName := "aws_apigatewayv2_api_route_responses"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-routes-routeid-routeresponses.html`,
		Resolver:    fetchApigatewayv2ApiRouteResponses,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   transformers.TransformWithStruct(&types.RouteResponse{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "api_route_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "route_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("route_id"),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   resolveApiRouteResponseArn,
				PrimaryKey: true,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2
	// No paginator available
	for {
		response, err := svc.GetRouteResponses(ctx, &config, func(options *apigatewayv2.Options) {
			options.Region = cl.Region
		})

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
