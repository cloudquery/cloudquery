package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Apigatewayv2Apis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_apis",
		Resolver:     fetchApigatewayv2Apis,
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
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "protocol_type",
				Type: schema.TypeString,
			},
			{
				Name: "route_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name: "api_endpoint",
				Type: schema.TypeString,
			},
			{
				Name: "api_gateway_managed",
				Type: schema.TypeBool,
			},
			{
				Name: "api_id",
				Type: schema.TypeString,
			},
			{
				Name: "api_key_selection_expression",
				Type: schema.TypeString,
			},
			{
				Name:     "cors_configuration_allow_credentials",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CorsConfiguration.AllowCredentials"),
			},
			{
				Name:     "cors_configuration_allow_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowHeaders"),
			},
			{
				Name:     "cors_configuration_allow_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowMethods"),
			},
			{
				Name:     "cors_configuration_allow_origins",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.AllowOrigins"),
			},
			{
				Name:     "cors_configuration_expose_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("CorsConfiguration.ExposeHeaders"),
			},
			{
				Name:     "cors_configuration_max_age",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CorsConfiguration.MaxAge"),
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
				Name: "disable_schema_validation",
				Type: schema.TypeBool,
			},
			{
				Name: "import_info",
				Type: schema.TypeStringArray,
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
				Name:     "aws_apigatewayv2_api_authorizers",
				Resolver: fetchApigatewayv2ApiAuthorizers,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "name",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_credentials_arn",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_id",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_payload_format_version",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_result_ttl_in_seconds",
						Type: schema.TypeInt,
					},
					{
						Name: "authorizer_type",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_uri",
						Type: schema.TypeString,
					},
					{
						Name: "enable_simple_responses",
						Type: schema.TypeBool,
					},
					{
						Name: "identity_source",
						Type: schema.TypeStringArray,
					},
					{
						Name: "identity_validation_expression",
						Type: schema.TypeString,
					},
					{
						Name:     "jwt_configuration_audience",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("JwtConfiguration.Audience"),
					},
					{
						Name:     "jwt_configuration_issuer",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("JwtConfiguration.Issuer"),
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_api_deployments",
				Resolver: fetchApigatewayv2ApiDeployments,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "auto_deployed",
						Type: schema.TypeBool,
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
						Name: "deployment_status",
						Type: schema.TypeString,
					},
					{
						Name: "deployment_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_api_integrations",
				Resolver: fetchApigatewayv2ApiIntegrations,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "connection_id",
						Type: schema.TypeString,
					},
					{
						Name: "connection_type",
						Type: schema.TypeString,
					},
					{
						Name: "content_handling_strategy",
						Type: schema.TypeString,
					},
					{
						Name: "credentials_arn",
						Type: schema.TypeString,
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "integration_id",
						Type: schema.TypeString,
					},
					{
						Name: "integration_method",
						Type: schema.TypeString,
					},
					{
						Name: "integration_response_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "integration_subtype",
						Type: schema.TypeString,
					},
					{
						Name: "integration_type",
						Type: schema.TypeString,
					},
					{
						Name: "integration_uri",
						Type: schema.TypeString,
					},
					{
						Name: "passthrough_behavior",
						Type: schema.TypeString,
					},
					{
						Name: "payload_format_version",
						Type: schema.TypeString,
					},
					{
						Name: "request_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "request_templates",
						Type: schema.TypeJSON,
					},
					{
						Name: "response_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "template_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "timeout_in_millis",
						Type: schema.TypeInt,
					},
					{
						Name:     "tls_config_server_name_to_verify",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("TlsConfig.ServerNameToVerify"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_apigatewayv2_api_integration_responses",
						Resolver: fetchApigatewayv2ApiIntegrationResponses,
						Columns: []schema.Column{
							{
								Name:     "api_integration_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "integration_response_key",
								Type: schema.TypeString,
							},
							{
								Name: "content_handling_strategy",
								Type: schema.TypeString,
							},
							{
								Name: "integration_response_id",
								Type: schema.TypeString,
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
								Name: "template_selection_expression",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_api_models",
				Resolver: fetchApigatewayv2ApiModels,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayv2apiModelModelTemplate,
					},
					{
						Name: "name",
						Type: schema.TypeString,
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
						Name: "model_id",
						Type: schema.TypeString,
					},
					{
						Name: "schema",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_api_routes",
				Resolver: fetchApigatewayv2ApiRoutes,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "route_key",
						Type: schema.TypeString,
					},
					{
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "api_key_required",
						Type: schema.TypeBool,
					},
					{
						Name: "authorization_scopes",
						Type: schema.TypeStringArray,
					},
					{
						Name: "authorization_type",
						Type: schema.TypeString,
					},
					{
						Name: "authorizer_id",
						Type: schema.TypeString,
					},
					{
						Name: "model_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "operation_name",
						Type: schema.TypeString,
					},
					{
						Name: "request_models",
						Type: schema.TypeJSON,
					},
					{
						Name: "request_parameters",
						Type: schema.TypeJSON,
					},
					{
						Name: "route_id",
						Type: schema.TypeString,
					},
					{
						Name: "route_response_selection_expression",
						Type: schema.TypeString,
					},
					{
						Name: "target",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_apigatewayv2_api_route_responses",
						Resolver: fetchApigatewayv2ApiRouteResponses,
						Columns: []schema.Column{
							{
								Name:     "api_route_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "route_response_key",
								Type: schema.TypeString,
							},
							{
								Name: "model_selection_expression",
								Type: schema.TypeString,
							},
							{
								Name: "response_models",
								Type: schema.TypeJSON,
							},
							{
								Name: "response_parameters",
								Type: schema.TypeJSON,
							},
							{
								Name: "route_response_id",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_apigatewayv2_api_stages",
				Resolver: fetchApigatewayv2ApiStages,
				Columns: []schema.Column{
					{
						Name:     "api_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "stage_name",
						Type: schema.TypeString,
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
						Name: "api_gateway_managed",
						Type: schema.TypeBool,
					},
					{
						Name: "auto_deploy",
						Type: schema.TypeBool,
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
						Name:     "route_settings_data_trace_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("DefaultRouteSettings.DataTraceEnabled"),
					},
					{
						Name:     "route_settings_detailed_metrics_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("DefaultRouteSettings.DetailedMetricsEnabled"),
					},
					{
						Name:     "route_settings_logging_level",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("DefaultRouteSettings.LoggingLevel"),
					},
					{
						Name:     "route_settings_throttling_burst_limit",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("DefaultRouteSettings.ThrottlingBurstLimit"),
					},
					{
						Name:     "route_settings_throttling_rate_limit",
						Type:     schema.TypeFloat,
						Resolver: schema.PathResolver("DefaultRouteSettings.ThrottlingRateLimit"),
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
						Name: "last_deployment_status_message",
						Type: schema.TypeString,
					},
					{
						Name: "last_updated_date",
						Type: schema.TypeTimestamp,
					},
					{
						Name: "route_settings",
						Type: schema.TypeJSON,
					},
					{
						Name: "stage_variables",
						Type: schema.TypeJSON,
					},
					{
						Name: "tags",
						Type: schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayv2Apis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config apigatewayv2.GetApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApis(ctx, &config, func(o *apigatewayv2.Options) {
			//o.Region = c.Region
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
func fetchApigatewayv2ApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetAuthorizersInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetAuthorizers(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetDeploymentsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDeployments(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiIntegrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetIntegrationsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetIntegrations(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Integration)
	if !ok {
		return fmt.Errorf("expected Integration but got %T", r)
	}
	p, ok := parent.Parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetIntegrationResponsesInput{
		ApiId:         p.ApiId,
		IntegrationId: r.IntegrationId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetIntegrationResponses(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetModelsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetModels(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func resolveApigatewayv2apiModelModelTemplate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.Model)
	if !ok {
		return fmt.Errorf("expected Model but got %T", r)
	}
	p, ok := resource.Parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetModelTemplateInput{
		ApiId:   p.ApiId,
		ModelId: r.ModelId,
	}
	client := meta.(*client.Client)
	svc := client.Services().Apigatewayv2

	response, err := svc.GetModelTemplate(ctx, &config, func(o *apigatewayv2.Options) {
		o.Region = client.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, response.Value)
}
func fetchApigatewayv2ApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected api but got %T", r)
	}
	config := apigatewayv2.GetRoutesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetRoutes(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Route)
	if !ok {
		return fmt.Errorf("expected Route but got %T", r)
	}
	p, ok := parent.Parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetRouteResponsesInput{
		ApiId:   p.ApiId,
		RouteId: r.RouteId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetRouteResponses(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
func fetchApigatewayv2ApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	r, ok := parent.Item.(types.Api)
	if !ok {
		return fmt.Errorf("expected Api but got %T", r)
	}
	config := apigatewayv2.GetStagesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetStages(ctx, &config, func(o *apigatewayv2.Options) {
			o.Region = c.Region
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
