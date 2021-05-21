package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ApigatewayRestApis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_rest_apis",
		Resolver:     fetchApigatewayRestApis,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "api_key_source",
				Type: schema.TypeString,
			},
			{
				Name: "binary_media_types",
				Type: schema.TypeStringArray,
			},
			{
				Name: "created_date",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "disable_execute_api_endpoint",
				Type: schema.TypeBool,
			},
			{
				Name:     "endpoint_configuration_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:     "endpoint_configuration_vpc_endpoint_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name: "minimum_compression_size",
				Type: schema.TypeInt,
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "policy",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name: "version",
				Type: schema.TypeString,
			},
			{
				Name: "warnings",
				Type: schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_apigateway_rest_api_authorizers",
				Resolver: fetchApigatewayRestApiAuthorizers,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "auth_type",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_credentials",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_result_ttl_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "authorizer_uri",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "identity_source",
						Type: schema.TypeString,
					},
					{
						Name: "identity_validation_expression",
						Type: schema.TypeString,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name:     "provider_arns",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("ProviderARNs"),
					},
					{
						Name: "type",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_deployments",
				Resolver: fetchApigatewayRestApiDeployments,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_summary",
						Type: schema.TypeJSON,
					},
					{
						Name: "created_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_documentation_parts",
				Resolver: fetchApigatewayRestApiDocumentationParts,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "documentation_part_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name:     "location_type",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Location.Type"),
					},
					{
						Name:     "location_method",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Location.Method"),
					},
					{
						Name:     "location_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Location.Name"),
					},
					{
						Name:     "location_path",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Location.Path"),
					},
					{
						Name:     "location_status_code",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Location.StatusCode"),
					},
					{
						Name: "properties",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_documentation_versions",
				Resolver: fetchApigatewayRestApiDocumentationVersions,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "created_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "version",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_gateway_responses",
				Resolver: fetchApigatewayRestApiGatewayResponses,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "default_response",
						Type: schema.TypeBool,
					},
					{
						Name: "response_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "response_templates",
						Type: schema.TypeJSON,
					},
					{
						Name: "response_type",
						Type: schema.TypeString,
					},
					{
						Name: "status_code",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_models",
				Resolver: fetchApigatewayRestApiModels,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayRestAPIModelModelTemplate,
					},
					{
						Name: "content_type",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "schema",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_request_validators",
				Resolver: fetchApigatewayRestApiRequestValidators,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "validate_request_body",
						Type: schema.TypeBool,
					},
					{
						Name: "validate_request_parameters",
						Type: schema.TypeBool,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_resources",
				Resolver: fetchApigatewayRestApiResources,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "resource_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "parent_id",
						Type: schema.TypeString,
					},
					{
						Name: "path",
						Type: schema.TypeString,
					},
					{
						Name: "path_part",
						Type: schema.TypeString,
					},
					{
						Name: "resource_methods",
						Type: schema.TypeJSON,
					},
				},
			},
			{
				Name:     "aws_apigateway_rest_api_stages",
				Resolver: fetchApigatewayRestApiStages,
				Columns: []schema.Column{
					{
						Name:     "rest_api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "access_log_settings_destination_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccessLogSettings.DestinationArn"),
					},
					{
						Name:     "access_log_settings_format",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("AccessLogSettings.Format"),
					},
					{
						Name: "cache_cluster_enabled",
						Type: schema.TypeBool,
					},
					{
						Name: "cache_cluster_size",
						Type: schema.TypeString,
					},
					{
						Name: "cache_cluster_status",
						Type: schema.TypeString,
					},
					{
						Name:     "canary_settings_deployment_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CanarySettings.DeploymentId"),
					},
					{
						Name:     "canary_settings_percent_traffic",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("CanarySettings.PercentTraffic"),
					},
					{
						Name:     "canary_settings_stage_variable_overrides",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("CanarySettings.StageVariableOverrides"),
					},
					{
						Name:     "canary_settings_use_stage_cache",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("CanarySettings.UseStageCache"),
					},
					{
						Name: "client_certificate_id",
						Type: schema.TypeString,
					},
					{
						Name: "created_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "deployment_id",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "documentation_version",
						Type: schema.TypeString,
					},
					{
						Name: "last_updated_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "method_settings",
						Type: schema.TypeJSON,
					},
					{
						Name: "stage_name",
						Type: schema.TypeString,
					},
					{
						Name: "tags",
						Type: schema.TypeJSON,
					},
					{
						Name: "tracing_enabled",
						Type: schema.TypeBool,
					},
					{
						Name: "variables",
						Type: schema.TypeJSON,
					},
					{
						Name: "web_acl_arn",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayRestApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigateway.GetRestApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetRestApis(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetAuthorizersInput{RestApiId: r.Id}
	for {
		response, err := svc.GetAuthorizers(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDeploymentsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDeployments(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiDocumentationParts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationPartsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationParts(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiDocumentationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationVersionsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationVersions(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiGatewayResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetGatewayResponsesInput{RestApiId: r.Id}
	for {
		response, err := svc.GetGatewayResponses(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetModelsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetModels(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func resolveApigatewayRestAPIModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.Model)
	if !ok {
		return fmt.Errorf("expected Model but got %T", r)
	}
	api, ok := resource.Parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	client := meta.(*client.Client)
	svc := client.Services().Apigateway

	config := apigateway.GetModelTemplateInput{
		RestApiId: api.Id,
		ModelName: r.Name,
	}

	response, err := svc.GetModelTemplate(ctx, &config, func(options *apigateway.Options) {
		options.Region = client.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Value)
}
func fetchApigatewayRestApiRequestValidators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetRequestValidatorsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetRequestValidators(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetResourcesInput{RestApiId: r.Id}
	for {
		response, err := svc.GetResources(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
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
func fetchApigatewayRestApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.RestApi)
	if !ok {
		return fmt.Errorf("expected RestApi but got %T", r)
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetStagesInput{RestApiId: r.Id}

	response, err := svc.GetStages(ctx, &config, func(options *apigateway.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.Item

	return nil
}
