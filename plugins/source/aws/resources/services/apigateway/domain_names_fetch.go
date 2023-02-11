package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayDomainNames(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apigateway.GetDomainNamesInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetDomainNamesPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayDomainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	domain := resource.Item.(types.DomainName)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/domainnames/%s", aws.ToString(domain.DomainName)),
	}.String())
}
func fetchApigatewayDomainNameBasePathMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.DomainName)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName, Limit: aws.Int32(500)}
	for p := apigateway.NewGetBasePathMappingsPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayDomainNameBasePathMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	domain := resource.Parent.Item.(types.DomainName)
	mapping := resource.Item.(types.BasePathMapping)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.ApigatewayService),
		Region:    cl.Region,
		AccountID: "",
		Resource:  fmt.Sprintf("/domainnames/%s/basepathmappings/%s", aws.ToString(domain.DomainName), aws.ToString(mapping.BasePath)),
	}.String())
}
