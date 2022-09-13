package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	apigatewayv2fix "github.com/cloudquery/cloudquery/plugins/source/aws/resources/forks/apigatewayv2"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayv2DomainNames(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var config apigatewayv2.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDomainNames(ctx, &config, func(options *apigatewayv2.Options) {
			options.Region = c.Region
			// NOTE: Swapping OperationDeserializer until this is fixed: https://github.com/aws/aws-sdk-go-v2/issues/1282
			options.APIOptions = append(options.APIOptions, apigatewayv2fix.SwapGetDomainNamesOperationDeserializer)
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
func resolveDomainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.DomainName)
	a := arn.ARN{
		Partition: cl.Partition,
		Service: "apigateway",
		Region: cl.Region,
		AccountID: cl.AccountID,
		Resource: "domainnames/" + aws.ToString(item.DomainName),
	}
	return resource.Set(c.Name, a.String())
}

func fetchApigatewayv2DomainNameApiMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.DomainName)
	config := apigatewayv2.GetApiMappingsInput{
		DomainName: r.DomainName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApiMappings(ctx, &config)

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


func resolveDomainNameApiMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ApiMapping)
	a := arn.ARN{
		Partition: cl.Partition,
		Service: "apigateway",
		Region: cl.Region,
		AccountID: cl.AccountID,
		Resource: "domainnames_api_mapping/" + aws.ToString(item.ApiMappingId),
	}
	return resource.Set(c.Name, a.String())
}