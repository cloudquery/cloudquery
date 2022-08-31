package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const domainNamesIDPart = "domainnames"

func ResolveApiModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Model)
	p := resource.Parent.Item.(types.Api)
	config := apigatewayv2.GetModelTemplateInput{
		ApiId:   p.ApiId,
		ModelId: r.ModelId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	response, err := svc.GetModelTemplate(ctx, &config)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, response.Value))
}

func ResolveDomainNameArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if err := client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		return []string{domainNamesIDPart, *resource.Item.(types.DomainName).DomainName}, nil
	})(ctx, meta, resource, c); err != nil {
		return diag.WrapError(err)
	}
	return nil
}

func ResolveApiMappingArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if err := client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.ApiMapping)
		p := resource.Parent.Item.(types.DomainName)
		return []string{domainNamesIDPart, *p.DomainName, "apimappings", *r.ApiMappingId}, nil
	})(ctx, meta, resource, c); err != nil {
		return diag.WrapError(err)
	}
	return nil
}

func ResolveVPCLinkArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if err := client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		return []string{"vpclinks", *resource.Item.(types.VpcLink).VpcLinkId}, nil
	})(ctx, meta, resource, c); err != nil {
		return diag.WrapError(err)
	}
	return nil
}
