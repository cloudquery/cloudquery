package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)


func fetchApigatewayApiKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := apigateway.GetApiKeysInput{
		IncludeValues: aws.Bool(true),
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	p := apigateway.NewGetApiKeysPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}


func resolveApiKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ApiKey)
	a := arn.ARN{
		Partition: cl.Partition,
		Service: "apigateway",
		Region: cl.Region,
		AccountID: cl.AccountID,
		Resource: "api_key/" + aws.ToString(item.Id),
	}
	return resource.Set(c.Name, a.String())
}