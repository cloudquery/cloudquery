package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	apiIDPart            = "/apis"
	apiRouteIDPart       = "routes"
	apiIntegrationIDPart = "integrations"
)

func fetchApigatewayv2Apis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigatewayv2.GetApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApis(ctx, &config)

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

func resolveApiArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Api)
	a := arn.ARN{
		Partition: cl.Partition,
		Service: "apigateway",
		Region: cl.Region,
		AccountID: cl.AccountID,
		Resource: "apis/" + aws.ToString(item.ApiId),
	}
	return resource.Set(c.Name, a.String())
}