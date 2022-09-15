package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApigatewayRestApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetRestApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for p := apigateway.NewGetRestApisPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	rapi := resource.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetAuthorizersInput{RestApiId: r.Id}
	for {
		response, err := svc.GetAuthorizers(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func resolveApigatewayRestAPIAuthorizerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	auth := resource.Item.(types.Authorizer)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "authorizers", *auth.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDeploymentsInput{RestApiId: r.Id}
	for p := apigateway.NewGetDeploymentsPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIDeploymentArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Item.(types.Deployment)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "deployments", *d.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiDocumentationParts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationPartsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationParts(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func resolveApigatewayRestAPIDocumentationPartArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Item.(types.DocumentationPart)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "documentation/parts", *d.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiDocumentationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationVersionsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationVersions(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func resolveApigatewayRestAPIDocumentationVersionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	v := resource.Item.(types.DocumentationVersion)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "documentation/versions", *v.Version)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiGatewayResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetGatewayResponsesInput{RestApiId: r.Id}
	for {
		response, err := svc.GetGatewayResponses(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func resolveApigatewayRestAPIGatewayResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.GatewayResponse)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "gatewayresponses", string(r.ResponseType))
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetModelsInput{RestApiId: r.Id}
	for p := apigateway.NewGetModelsPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIModelArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	m := resource.Item.(types.Model)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "models", *m.Name)
	return resource.Set(c.Name, arn)
}
func resolveApigatewayRestAPIModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Model)
	api := resource.Parent.Item.(types.RestApi)
	cl := meta.(*client.Client)
	svc := cl.Services().Apigateway

	if api.Id == nil || r.Name == nil {
		return nil
	}

	config := apigateway.GetModelTemplateInput{
		RestApiId: api.Id,
		ModelName: r.Name,
	}

	response, err := svc.GetModelTemplate(ctx, &config)
	if err != nil {
		if client.IsAWSError(err, "BadRequestException") {
			// This is an application level error and the user has nothing to do with that.
			// https://github.com/cloudquery/cq-provider-aws/pull/567#discussion_r827095787
			// The suer will be able to find incorrect configured models via
			// select * from aws_apigateway_rest_api_models where model_template is nil
			return nil
		}
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, response.Value)
}
func fetchApigatewayRestApiRequestValidators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetRequestValidatorsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetRequestValidators(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func resolveApigatewayRestAPIRequestValidatorArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.RequestValidator)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "requestvalidators", *r.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetResourcesInput{RestApiId: r.Id}
	for p := apigateway.NewGetResourcesPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIResourceArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.Resource)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "resources", *r.Id)
	return resource.Set(c.Name, arn)
}
func fetchApigatewayRestApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetStagesInput{RestApiId: r.Id}

	response, err := svc.GetStages(ctx, &config)
	if err != nil {
		if c.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- response.Item

	return nil
}
func resolveApigatewayRestAPIStageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	s := resource.Item.(types.Stage)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "stages", *s.StageName)
	return resource.Set(c.Name, arn)
}
