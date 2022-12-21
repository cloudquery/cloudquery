package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

const (
	apiIDPart            = "/apis"
	apiRouteIDPart       = "routes"
	apiIntegrationIDPart = "integrations"
)

func resolveApiArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		return []string{apiIDPart, *resource.Item.(types.Api).ApiId}, nil
	})
}
func resolveApiAuthorizerArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Authorizer)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, "authorizers", *r.AuthorizerId}, nil
	})
}
func resolveApiDeploymentArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Deployment)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, "deployments", *r.DeploymentId}, nil
	})
}
func resolveApiIntegrationArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Integration)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, apiIntegrationIDPart, *r.IntegrationId}, nil
	})
}
func resolveApiIntegrationResponseArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.IntegrationResponse)
		i := resource.Parent.Item.(types.Integration)
		api := resource.Parent.Parent.Item.(types.Api)
		return []string{apiIDPart, *api.ApiId, apiIntegrationIDPart, *i.IntegrationId, "integrationresponses", *r.IntegrationResponseId}, nil
	})
}
func resolveApiModelArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Model)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, "models", *r.ModelId}, nil
	})
}
func resolveApiRouteArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Route)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, apiRouteIDPart, *r.RouteId}, nil
	})
}
func resolveApiRouteResponseArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.RouteResponse)
		route := resource.Parent.Item.(types.Route)
		api := resource.Parent.Parent.Item.(types.Api)
		return []string{apiIDPart, *api.ApiId, apiRouteIDPart, *route.RouteId, "routeresponses", *r.RouteResponseId}, nil
	})
}
func resolveApiStageArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		r := resource.Item.(types.Stage)
		p := resource.Parent.Item.(types.Api)
		return []string{apiIDPart, *p.ApiId, "stages", *r.StageName}, nil
	})
}
func resolveVpcLinkArn() schema.ColumnResolver {
	return client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
		return []string{"vpclinks", *resource.Item.(types.VpcLink).VpcLinkId}, nil
	})
}
func fetchApigatewayv2Apis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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
func fetchApigatewayv2ApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetAuthorizersInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetAuthorizers(ctx, &config)

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
func fetchApigatewayv2ApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetDeploymentsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDeployments(ctx, &config)

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
func fetchApigatewayv2ApiIntegrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetIntegrationsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetIntegrations(ctx, &config)

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
func fetchApigatewayv2ApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Integration)
	p := parent.Parent.Item.(types.Api)
	config := apigatewayv2.GetIntegrationResponsesInput{
		ApiId:         p.ApiId,
		IntegrationId: r.IntegrationId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetIntegrationResponses(ctx, &config)

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
func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetModelsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetModels(ctx, &config)

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
func resolveApigatewayv2apiModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
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
		return err
	}
	return resource.Set(c.Name, response.Value)
}
func fetchApigatewayv2ApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetRoutesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetRoutes(ctx, &config)

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
func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Route)
	p := parent.Parent.Item.(types.Api)
	config := apigatewayv2.GetRouteResponsesInput{
		ApiId:   p.ApiId,
		RouteId: r.RouteId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
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
func fetchApigatewayv2ApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetStagesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetStages(ctx, &config)

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
