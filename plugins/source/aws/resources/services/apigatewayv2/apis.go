package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource apis --config gen.hcl --output .
func Apis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_apis",
		Description:  "Represents an API",
		Resolver:     fetchApigatewayv2Apis,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayv2apiArn,
			},
			{
				Name:        "name",
				Description: "The name of the API",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol_type",
				Description: "The API protocol",
				Type:        schema.TypeString,
			},
			{
				Name:        "route_selection_expression",
				Description: "The route selection expression for the API",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_endpoint",
				Description: "The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_gateway_managed",
				Description: "Specifies whether an API is managed by API Gateway",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The API ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApiId"),
			},
			{
				Name:        "api_key_selection_expression",
				Description: "An API key selection expression",
				Type:        schema.TypeString,
			},
			{
				Name:        "cors_configuration_allow_credentials",
				Description: "Specifies whether credentials are included in the CORS request",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowCredentials"),
			},
			{
				Name:        "cors_configuration_allow_headers",
				Description: "Represents a collection of allowed headers",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowHeaders"),
			},
			{
				Name:        "cors_configuration_allow_methods",
				Description: "Represents a collection of allowed HTTP methods",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowMethods"),
			},
			{
				Name:        "cors_configuration_allow_origins",
				Description: "Represents a collection of allowed origins",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowOrigins"),
			},
			{
				Name:        "cors_configuration_expose_headers",
				Description: "Represents a collection of exposed headers",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.ExposeHeaders"),
			},
			{
				Name:        "cors_configuration_max_age",
				Description: "The number of seconds that the browser should cache preflight request results",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("CorsConfiguration.MaxAge"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The description of the API",
				Type:        schema.TypeString,
			},
			{
				Name:        "disable_execute_api_endpoint",
				Description: "Specifies whether clients can invoke your API by using the default execute-api endpoint",
				Type:        schema.TypeBool,
			},
			{
				Name:        "disable_schema_validation",
				Description: "Avoid validating models when creating a deployment",
				Type:        schema.TypeBool,
			},
			{
				Name:        "import_info",
				Description: "The validation information during API import",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags",
				Description: "A collection of tags associated with the API",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "version",
				Description: "A version identifier for the API",
				Type:        schema.TypeString,
			},
			{
				Name:        "warnings",
				Description: "The warning messages reported when failonwarnings is turned on during API import",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_api_authorizers",
				Description: "Represents an authorizer",
				Resolver:    fetchApigatewayv2ApiAuthorizers,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiAuthorizerArn,
					},
					{
						Name:        "name",
						Description: "The name of the authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_credentials_arn",
						Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_id",
						Description: "The authorizer identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_payload_format_version",
						Description: "Specifies the format of the payload sent to an HTTP API Lambda authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_result_ttl_in_seconds",
						Description: "The time to live (TTL) for cached authorizer results, in seconds",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "authorizer_type",
						Description: "The authorizer type",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_uri",
						Description: "The authorizer's Uniform Resource Identifier (URI)",
						Type:        schema.TypeString,
					},
					{
						Name:        "enable_simple_responses",
						Description: "Specifies whether a Lambda authorizer returns a response in a simple format",
						Type:        schema.TypeBool,
					},
					{
						Name:        "identity_source",
						Description: "The identity source for which authorization is requested",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "identity_validation_expression",
						Description: "The validation expression does not apply to the REQUEST authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "jwt_configuration_audience",
						Description: "A list of the intended recipients of the JWT",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("JwtConfiguration.Audience"),
					},
					{
						Name:        "jwt_configuration_issuer",
						Description: "The base domain of the identity provider that issues JSON Web Tokens",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("JwtConfiguration.Issuer"),
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_deployments",
				Description: "An immutable representation of an API that can be called by users",
				Resolver:    fetchApigatewayv2ApiDeployments,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiDeploymentArn,
					},
					{
						Name:        "auto_deployed",
						Description: "Specifies whether a deployment was automatically released",
						Type:        schema.TypeBool,
					},
					{
						Name:        "created_date",
						Description: "The date and time when the Deployment resource was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "deployment_id",
						Description: "The identifier for the deployment",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_status",
						Description: "The status of the deployment: PENDING, FAILED, or SUCCEEDED",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_status_message",
						Description: "May contain additional feedback on the status of an API deployment",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description for the deployment",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_integrations",
				Description: "Represents an integration",
				Resolver:    fetchApigatewayv2ApiIntegrations,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiIntegrationArn,
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether an integration is managed by API Gateway",
						Type:        schema.TypeBool,
					},
					{
						Name:        "connection_id",
						Description: "The ID of the VPC link for a private integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_type",
						Description: "The type of the network connection to the integration endpoint",
						Type:        schema.TypeString,
					},
					{
						Name:        "content_handling_strategy",
						Description: "Supported only for WebSocket APIs",
						Type:        schema.TypeString,
					},
					{
						Name:        "credentials_arn",
						Description: "Specifies the credentials required for the integration, if any",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "Represents the description of an integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_id",
						Description: "Represents the identifier of an integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_method",
						Description: "Specifies the integration's HTTP method type",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_response_selection_expression",
						Description: "The integration response selection expression for the integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_subtype",
						Description: "Supported only for HTTP API AWS_PROXY integrations",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_type",
						Description: "The integration type of an integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_uri",
						Description: "For a Lambda integration, specify the URI of a Lambda function",
						Type:        schema.TypeString,
					},
					{
						Name:        "passthrough_behavior",
						Description: "Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "payload_format_version",
						Description: "Specifies the format of the payload sent to an integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "request_parameters",
						Description: "For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "request_templates",
						Description: "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "response_parameters",
						Description: "Supported only for HTTP APIs",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "template_selection_expression",
						Description: "The template selection expression for the integration",
						Type:        schema.TypeString,
					},
					{
						Name:        "timeout_in_millis",
						Description: "Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "tls_config_server_name_to_verify",
						Description: "If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TlsConfig.ServerNameToVerify"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_apigatewayv2_api_integration_responses",
						Description: "Represents an integration response",
						Resolver:    fetchApigatewayv2ApiIntegrationResponses,
						Columns: []schema.Column{
							{
								Name:        "api_integration_cq_id",
								Description: "Unique CloudQuery ID of aws_apigatewayv2_api_integrations table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "integration_id",
								Description: "Represents the identifier of an integration",
								Type:        schema.TypeString,
								Resolver:    schema.ParentPathResolver("IntegrationId"),
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) for the resource",
								Type:        schema.TypeString,
								Resolver:    resolveApigatewayv2apiIntegrationResponseArn,
							},
							{
								Name:        "integration_response_key",
								Description: "The integration response key",
								Type:        schema.TypeString,
							},
							{
								Name:        "content_handling_strategy",
								Description: "Supported only for WebSocket APIs",
								Type:        schema.TypeString,
							},
							{
								Name:        "integration_response_id",
								Description: "The integration response ID",
								Type:        schema.TypeString,
							},
							{
								Name:        "response_parameters",
								Description: "A key-value map specifying response parameters that are passed to the method response from the backend",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "response_templates",
								Description: "The collection of response templates for the integration response as a string-to-string map of key-value pairs",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "template_selection_expression",
								Description: "The template selection expressions for the integration response",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_models",
				Description: "Represents a data model for an API",
				Resolver:    fetchApigatewayv2ApiModels,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiModelArn,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayv2apiModelModelTemplate,
					},
					{
						Name:        "name",
						Description: "The name of the model",
						Type:        schema.TypeString,
					},
					{
						Name:        "content_type",
						Description: "The content-type for the model, for example, \"application/json\"",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the model",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_id",
						Description: "The model identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "schema",
						Description: "The schema for the model",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_routes",
				Description: "Represents a route",
				Resolver:    fetchApigatewayv2ApiRoutes,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiRouteArn,
					},
					{
						Name:        "route_key",
						Description: "The route key for the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether a route is managed by API Gateway",
						Type:        schema.TypeBool,
					},
					{
						Name:        "api_key_required",
						Description: "Specifies whether an API key is required for this route",
						Type:        schema.TypeBool,
					},
					{
						Name:        "authorization_scopes",
						Description: "A list of authorization scopes configured on a route",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "authorization_type",
						Description: "The authorization type for the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_id",
						Description: "The identifier of the Authorizer resource to be associated with this route",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_selection_expression",
						Description: "The model selection expression for the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "operation_name",
						Description: "The operation name for the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "request_models",
						Description: "The request models for the route",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "request_parameters",
						Description: "The request parameters for the route",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "route_id",
						Description: "The route ID",
						Type:        schema.TypeString,
					},
					{
						Name:        "route_response_selection_expression",
						Description: "The route response selection expression for the route",
						Type:        schema.TypeString,
					},
					{
						Name:        "target",
						Description: "The target for the route",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_apigatewayv2_api_route_responses",
						Description: "Represents a route response",
						Resolver:    fetchApigatewayv2ApiRouteResponses,
						Columns: []schema.Column{
							{
								Name:        "api_route_cq_id",
								Description: "Unique CloudQuery ID of aws_apigatewayv2_api_routes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "route_id",
								Description: "The Route id",
								Type:        schema.TypeString,
								Resolver:    schema.ParentPathResolver("RouteId"),
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) for the resource",
								Type:        schema.TypeString,
								Resolver:    resolveApigatewayv2apiRouteResponseArn,
							},
							{
								Name:        "route_response_key",
								Description: "Represents the route response key of a route response",
								Type:        schema.TypeString,
							},
							{
								Name:        "model_selection_expression",
								Description: "Represents the model selection expression of a route response",
								Type:        schema.TypeString,
							},
							{
								Name:        "response_models",
								Description: "Represents the response models of a route response",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "response_parameters",
								Description: "Represents the response parameters of a route response",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "route_response_id",
								Description: "Represents the identifier of a route response",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_stages",
				Description: "Represents an API stage",
				Resolver:    fetchApigatewayv2ApiStages,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API id",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("ApiId"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayv2apiStageArn,
					},
					{
						Name:        "stage_name",
						Description: "The name of the stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_log_settings_destination_arn",
						Description: "The ARN of the CloudWatch Logs log group to receive access logs",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.DestinationArn"),
					},
					{
						Name:        "access_log_settings_format",
						Description: "A single line format of the access logs of data, as specified by selected $context variables",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.Format"),
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether a stage is managed by API Gateway",
						Type:        schema.TypeBool,
					},
					{
						Name:        "auto_deploy",
						Description: "Specifies whether updates to an API automatically trigger a new deployment",
						Type:        schema.TypeBool,
					},
					{
						Name:        "client_certificate_id",
						Description: "The identifier of a client certificate for a Stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_date",
						Description: "The timestamp when the stage was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "route_settings_data_trace_enabled",
						Description: "Specifies whether (true) or not (false) data trace logging is enabled for this route",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultRouteSettings.DataTraceEnabled"),
					},
					{
						Name:        "route_settings_detailed_metrics_enabled",
						Description: "Specifies whether detailed metrics are enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultRouteSettings.DetailedMetricsEnabled"),
					},
					{
						Name:        "route_settings_logging_level",
						Description: "Specifies the logging level for this route: INFO, ERROR, or OFF",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultRouteSettings.LoggingLevel"),
					},
					{
						Name:        "route_settings_throttling_burst_limit",
						Description: "Specifies the throttling burst limit",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DefaultRouteSettings.ThrottlingBurstLimit"),
					},
					{
						Name:        "route_settings_throttling_rate_limit",
						Description: "Specifies the throttling rate limit",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DefaultRouteSettings.ThrottlingRateLimit"),
					},
					{
						Name:        "deployment_id",
						Description: "The identifier of the Deployment that the Stage is associated with",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_deployment_status_message",
						Description: "Describes the status of the last deployment of a stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_updated_date",
						Description: "The timestamp when the stage was last updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "route_settings",
						Description: "Route settings for the stage, by routeKey",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "stage_variables",
						Description: "A map that defines the stage variables for a stage resource",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "tags",
						Description: "The collection of tags",
						Type:        schema.TypeJSON,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchApigatewayv2Apis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigatewayv2.GetApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetApis(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Item.(types.Api)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetAuthorizersInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetAuthorizers(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiAuthorizerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	auth := resource.Item.(types.Authorizer)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, "authorizers", *auth.AuthorizerId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetDeploymentsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetDeployments(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiDeploymentArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	d := resource.Item.(types.Deployment)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, "deployments", *d.DeploymentId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiIntegrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetIntegrationsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetIntegrations(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiIntegrationArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	int := resource.Item.(types.Integration)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, apiIntegrationIDPart, *int.IntegrationId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiIntegrationResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Parent.Item.(types.Api)
	int := resource.Parent.Item.(types.Integration)
	r := resource.Item.(types.IntegrationResponse)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, apiIntegrationIDPart, *int.IntegrationId, "integrationresponses", *r.IntegrationResponseId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetModelsInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetModels(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiModelArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	m := resource.Item.(types.Model)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, "models", *m.ModelId)
	return diag.WrapError(resource.Set(c.Name, arn))
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, response.Value))
}
func fetchApigatewayv2ApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetRoutesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetRoutes(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiRouteArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	r := resource.Item.(types.Route)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, apiRouteIDPart, *r.RouteId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiRouteResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Parent.Item.(types.Api)
	route := resource.Parent.Item.(types.Route)
	resp := resource.Item.(types.RouteResponse)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, apiRouteIDPart, *route.RouteId, "routeresponses", *resp.RouteResponseId)
	return diag.WrapError(resource.Set(c.Name, arn))
}
func fetchApigatewayv2ApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Api)
	config := apigatewayv2.GetStagesInput{
		ApiId: r.ApiId,
	}
	c := meta.(*client.Client)
	svc := c.Services().Apigatewayv2
	for {
		response, err := svc.GetStages(ctx, &config)

		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func resolveApigatewayv2apiStageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	api := resource.Parent.Item.(types.Api)
	s := resource.Item.(types.Stage)
	arn := cl.RegionGlobalARN(client.ApigatewayService, apiIDPart, *api.ApiId, "stages", *s.StageName)
	return diag.WrapError(resource.Set(c.Name, arn))
}
