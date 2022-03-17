package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const (
	apiIDPart            = "/apis"
	apiRouteIDPart       = "routes"
	apiIntegrationIDPart = "integrations"
)

func Apigatewayv2Apis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigatewayv2_apis",
		Description:  "Represents an API.",
		Resolver:     fetchApigatewayv2Apis,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
					return []string{apiIDPart, *resource.Item.(types.Api).ApiId}, nil
				}),
			},
			{
				Name:        "name",
				Description: "The name of the API.",
				Type:        schema.TypeString,
			},
			{
				Name:        "protocol_type",
				Description: "The API protocol.",
				Type:        schema.TypeString,
			},
			{
				Name:        "route_selection_expression",
				Description: "The route selection expression for the API. For HTTP APIs, the routeSelectionExpression must be ${request.method} ${request.path}. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs.",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_endpoint",
				Description: "The URI of the API, of the form {api-id}.execute-api.{region}.amazonaws.com. The stage name is typically appended to this URI to form a complete path to a deployed API stage.",
				Type:        schema.TypeString,
			},
			{
				Name:        "api_gateway_managed",
				Description: "Specifies whether an API is managed by API Gateway. You can't update or delete a managed API by using API Gateway. A managed API can be deleted only through the tooling or service that created it.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "id",
				Description: "The API ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApiId"),
			},
			{
				Name:        "api_key_selection_expression",
				Description: "An API key selection expression. Supported only for WebSocket APIs. See API Key Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions).",
				Type:        schema.TypeString,
			},
			{
				Name:        "cors_configuration_allow_credentials",
				Description: "Specifies whether credentials are included in the CORS request. Supported only for HTTP APIs.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowCredentials"),
			},
			{
				Name:        "cors_configuration_allow_headers",
				Description: "Represents a collection of allowed headers. Supported only for HTTP APIs.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowHeaders"),
			},
			{
				Name:        "cors_configuration_allow_methods",
				Description: "Represents a collection of allowed HTTP methods. Supported only for HTTP APIs.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowMethods"),
			},
			{
				Name:        "cors_configuration_allow_origins",
				Description: "Represents a collection of allowed origins. Supported only for HTTP APIs.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("CorsConfiguration.AllowOrigins"),
			},
			{
				Name:          "cors_configuration_expose_headers",
				Description:   "Represents a collection of exposed headers. Supported only for HTTP APIs.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("CorsConfiguration.ExposeHeaders"),
				IgnoreInTests: true,
			},
			{
				Name:        "cors_configuration_max_age",
				Description: "The number of seconds that the browser should cache preflight request results. Supported only for HTTP APIs.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CorsConfiguration.MaxAge"),
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "description",
				Description:   "The description of the API.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "disable_execute_api_endpoint",
				Description: "Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "disable_schema_validation",
				Description: "Avoid validating models when creating a deployment. Supported only for WebSocket APIs.",
				Type:        schema.TypeBool,
			},
			{
				Name:          "import_info",
				Description:   "The validation information during API import. This may include particular properties of your OpenAPI definition which are ignored during import. Supported only for HTTP APIs.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "A collection of tags associated with the API.",
				Type:        schema.TypeJSON,
			},
			{
				Name:          "version",
				Description:   "A version identifier for the API.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "warnings",
				Description:   "The warning messages reported when failonwarnings is turned on during API import.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_apigatewayv2_api_authorizers",
				Description: "Represents an authorizer.",
				Resolver:    fetchApigatewayv2ApiAuthorizers,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "authorizer_id"}},
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Authorizer)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, "authorizers", *r.AuthorizerId}, nil
						}),
					},
					{
						Name:        "name",
						Description: "The name of the authorizer.",
						Type:        schema.TypeString,
					},
					{
						Name:          "authorizer_credentials_arn",
						Description:   "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, don't specify this parameter. Supported only for REQUEST authorizers.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "authorizer_id",
						Description: "The authorizer identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:          "authorizer_payload_format_version",
						Description:   "Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are 1.0 and 2.0. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html).",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "authorizer_result_ttl_in_seconds",
						Description: "The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "authorizer_type",
						Description: "The authorizer type. Specify REQUEST for a Lambda function using incoming request parameters. Specify JWT to use JSON Web Tokens (supported only for HTTP APIs).",
						Type:        schema.TypeString,
					},
					{
						Name:          "authorizer_uri",
						Description:   "The authorizer's Uniform Resource Identifier (URI). For REQUEST authorizers, this must be a well-formed Lambda function URI, for example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations. In general, the URI has this form: arn:aws:apigateway:{region}:lambda:path/{service_api} , where {region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial /. For Lambda functions, this is usually of the form /2015-03-31/functions/[FunctionARN]/invocations. Supported only for REQUEST authorizers.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "enable_simple_responses",
						Description: "Specifies whether a Lambda authorizer returns a response in a simple format. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "identity_source",
						Description: "The identity source for which authorization is requested. For a REQUEST authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with $, for example, $request.header.Auth, $request.querystring.Name. These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see Working with AWS Lambda authorizers for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html). For JWT, a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example $request.header.Authorization.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:          "identity_validation_expression",
						Description:   "The validation expression does not apply to the REQUEST authorizer.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "jwt_configuration_audience",
						Description: "A list of the intended recipients of the JWT. A valid JWT must provide an aud that matches at least one entry in this list. See RFC 7519 (https://tools.ietf.org/html/rfc7519#section-4.1.3). Supported only for HTTP APIs.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("JwtConfiguration.Audience"),
					},
					{
						Name:        "jwt_configuration_issuer",
						Description: "The base domain of the identity provider that issues JSON Web Tokens. For example, an Amazon Cognito user pool has the following format: https://cognito-idp.{region}.amazonaws.com/{userPoolId} . Required for the JWT authorizer type. Supported only for HTTP APIs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("JwtConfiguration.Issuer"),
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_deployments",
				Description: "An immutable representation of an API that can be called by users.",
				Resolver:    fetchApigatewayv2ApiDeployments,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "deployment_id"}},
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Deployment)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, "deployments", *r.DeploymentId}, nil
						}),
					},
					{
						Name:        "auto_deployed",
						Description: "Specifies whether a deployment was automatically released.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "created_date",
						Description: "The date and time when the Deployment resource was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "deployment_id",
						Description: "The identifier for the deployment.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deployment_status",
						Description: "The status of the deployment: PENDING, FAILED, or SUCCEEDED.",
						Type:        schema.TypeString,
					},
					{
						Name:          "deployment_status_message",
						Description:   "May contain additional feedback on the status of an API deployment.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "description",
						Description: "The description for the deployment.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_integrations",
				Description: "Represents an integration.",
				Resolver:    fetchApigatewayv2ApiIntegrations,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "integration_id"}},
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Integration)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, apiIntegrationIDPart, *r.IntegrationId}, nil
						}),
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether an integration is managed by API Gateway. If you created an API using using quick create, the resulting integration is managed by API Gateway. You can update a managed integration, but you can't delete it.",
						Type:        schema.TypeBool,
					},
					{
						Name:          "connection_id",
						Description:   "The ID of the VPC link for a private integration. Supported only for HTTP APIs.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "connection_type",
						Description: "The type of the network connection to the integration endpoint. Specify INTERNET for connections through the public routable internet or VPC_LINK for private connections between API Gateway and resources in a VPC. The default value is INTERNET.",
						Type:        schema.TypeString,
					},
					{
						Name:        "content_handling_strategy",
						Description: "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded string. If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.",
						Type:        schema.TypeString,
					},
					{
						Name:        "credentials_arn",
						Description: "Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, specify null.",
						Type:        schema.TypeString,
					},
					{
						Name:          "description",
						Description:   "Represents the description of an integration.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "integration_id",
						Description: "Represents the identifier of an integration.",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_method",
						Description: "Specifies the integration's HTTP method type.",
						Type:        schema.TypeString,
					},
					{
						Name:          "integration_response_selection_expression",
						Description:   "The integration response selection expression for the integration. Supported only for WebSocket APIs. See Integration Response Selection Expressions (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-integration-response-selection-expressions).",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "integration_subtype",
						Description: "Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service action to invoke. To learn more, see Integration subtype reference (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services-reference.html).",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_type",
						Description: "The integration type of an integration. One of the following: AWS: for integrating the route or method request with an AWS service action, including the Lambda function-invoking action. With the Lambda function-invoking action, this is referred to as the Lambda custom integration. With any other AWS service action, this is known as AWS integration. Supported only for WebSocket APIs. AWS_PROXY: for integrating the route or method request with a Lambda function or other AWS service action. This integration is also referred to as a Lambda proxy integration. HTTP: for integrating the route or method request with an HTTP endpoint. This integration is also referred to as the HTTP custom integration. Supported only for WebSocket APIs. HTTP_PROXY: for integrating the route or method request with an HTTP endpoint, with the client request passed through as-is. This is also referred to as HTTP proxy integration. MOCK: for integrating the route or method request with API Gateway as a \"loopback\" endpoint without invoking any backend. Supported only for WebSocket APIs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "integration_uri",
						Description: "For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. If you specify the ARN of an AWS Cloud Map service, API Gateway uses DiscoverInstances to identify resources. You can use query parameters to target specific resources. To learn more, see DiscoverInstances (https://docs.aws.amazon.com/cloud-map/latest/api/API_DiscoverInstances.html). For private integrations, all resources must be owned by the same AWS account.",
						Type:        schema.TypeString,
					},
					{
						Name:        "passthrough_behavior",
						Description: "Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource. There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported only for WebSocket APIs. WHEN_NO_MATCH passes the request body for unmapped content types through to the integration backend without transformation. NEVER rejects unmapped content types with an HTTP 415 Unsupported Media Type response. WHEN_NO_TEMPLATES allows pass-through when the integration has no content types mapped to templates. However, if there is at least one content type defined, unmapped content types will be rejected with the same HTTP 415 Unsupported Media Type response.",
						Type:        schema.TypeString,
					},
					{
						Name:        "payload_format_version",
						Description: "Specifies the format of the payload sent to an integration. Required for HTTP APIs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "request_parameters",
						Description: "For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend. The key is an integration request parameter name and the associated value is a method request parameter value or static value that must be enclosed within single quotes and pre-encoded as required by the backend. The method request parameter value must match the pattern of method.request.{location}.{name} , where {location} is querystring, path, or header; and {name} must be a valid and unique method request parameter name. For HTTP API integrations with a specified integrationSubtype, request parameters are a key-value map specifying parameters that are passed to AWS_PROXY integrations. You can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Working with AWS service integrations for HTTP APIs (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-develop-integrations-aws-services.html). For HTTP API itegrations, without a specified integrationSubtype request parameters are a key-value map specifying how to transform HTTP requests before sending them to backend integrations. The key should follow the pattern <action>:<header|querystring|path>.<location>. The action can be append, overwrite or remove. For values, you can provide static values, or map request data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).",
						Type:        schema.TypeJSON,
					},
					{
						Name:          "request_templates",
						Description:   "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. The content type value is the key in this map, and the template (as a String) is the value. Supported only for WebSocket APIs.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:        "response_parameters",
						Description: "Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend integration before returning the response to clients. Specify a key-value map from a selection key to response parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key must match pattern <action>:<header>.<location> or overwrite.statuscode. The action can be append, overwrite or remove. The value can be a static value, or map to response data, stage variables, or context variables that are evaluated at runtime. To learn more, see Transforming API requests and responses (https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-parameter-mapping.html).",
						Type:        schema.TypeJSON,
					},
					{
						Name:          "template_selection_expression",
						Description:   "The template selection expression for the integration. Supported only for WebSocket APIs.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "timeout_in_millis",
						Description: "Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.",
						Type:        schema.TypeInt,
					},
					{
						Name:          "tls_config_server_name_to_verify",
						Description:   "If you specify a server name, API Gateway uses it to verify the hostname on the integration's certificate. The server name is also included in the TLS handshake to support Server Name Indication (SNI) or virtual hosting.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("TlsConfig.ServerNameToVerify"),
						IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_apigatewayv2_api_integration_responses",
						Description:   "Represents an integration response.",
						Resolver:      fetchApigatewayv2ApiIntegrationResponses,
						Options:       schema.TableCreationOptions{PrimaryKeys: []string{"api_integration_cq_id", "integration_response_id"}},
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "api_integration_cq_id",
								Description: "Unique CloudQuery ID of aws_apigatewayv2_api_integrations table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "integration_id",
								Description: "Represents the identifier of an integration.",
								Type:        schema.TypeString,
								Resolver:    schema.ParentResourceFieldResolver("integration_id"),
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) for the resource.",
								Type:        schema.TypeString,
								Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
									r := resource.Item.(types.IntegrationResponse)
									i := resource.Parent.Item.(types.Integration)
									api := resource.Parent.Parent.Item.(types.Api)
									return []string{apiIDPart, *api.ApiId, apiIntegrationIDPart, *i.IntegrationId, "integrationresponses", *r.IntegrationResponseId}, nil
								}),
							},
							{
								Name:        "integration_response_key",
								Description: "The integration response key.",
								Type:        schema.TypeString,
							},
							{
								Name:        "content_handling_strategy",
								Description: "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors: CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded string to the corresponding binary blob. CONVERT_TO_TEXT: Converts a response payload from a binary blob to a Base64-encoded string. If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.",
								Type:        schema.TypeString,
							},
							{
								Name:        "integration_response_id",
								Description: "The integration response ID.",
								Type:        schema.TypeString,
							},
							{
								Name:        "response_parameters",
								Description: "A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of method.response.header.{name}, where name is a valid and unique header name. The mapped non-static value must match the pattern of integration.response.header.{name} or integration.response.body.{JSON-expression}, where name is a valid and unique response header name and JSON-expression is a valid JSON expression without the $ prefix.",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "response_templates",
								Description: "The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "template_selection_expression",
								Description: "The template selection expressions for the integration response.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:          "aws_apigatewayv2_api_models",
				Description:   "Represents a data model for an API.",
				Resolver:      fetchApigatewayv2ApiModels,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "model_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Model)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, "models", *r.ModelId}, nil
						}),
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayv2apiModelModelTemplate,
					},
					{
						Name:        "name",
						Description: "The name of the model. Must be alphanumeric.",
						Type:        schema.TypeString,
					},
					{
						Name:        "content_type",
						Description: "The content-type for the model, for example, \"application/json\".",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the model.",
						Type:        schema.TypeString,
					},
					{
						Name:        "model_id",
						Description: "The model identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "schema",
						Description: "The schema for the model. For application/json models, this should be JSON schema draft 4 model.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_routes",
				Description: "Represents a route.",
				Resolver:    fetchApigatewayv2ApiRoutes,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "route_id"}},
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Route)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, apiRouteIDPart, *r.RouteId}, nil
						}),
					},
					{
						Name:        "route_key",
						Description: "The route key for the route.",
						Type:        schema.TypeString,
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether a route is managed by API Gateway. If you created an API using quick create, the $default route is managed by API Gateway. You can't modify the $default route key.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "api_key_required",
						Description: "Specifies whether an API key is required for this route. Supported only for WebSocket APIs.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "authorization_scopes",
						Description: "A list of authorization scopes configured on a route. The scopes are used with a JWT authorizer to authorize the method invocation. The authorization works by matching the route scopes against the scopes parsed from the access token in the incoming request. The method invocation is authorized if any route scope matches a claimed scope in the access token. Otherwise, the invocation is not authorized. When the route scope is configured, the client must provide an access token instead of an identity token for authorization purposes.",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "authorization_type",
						Description: "The authorization type for the route. For WebSocket APIs, valid values are NONE for open access, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer For HTTP APIs, valid values are NONE for open access, JWT for using JSON Web Tokens, AWS_IAM for using AWS IAM permissions, and CUSTOM for using a Lambda authorizer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_id",
						Description: "The identifier of the Authorizer resource to be associated with this route. The authorizer identifier is generated by API Gateway when you created the authorizer.",
						Type:        schema.TypeString,
					},
					{
						Name:          "model_selection_expression",
						Description:   "The model selection expression for the route. Supported only for WebSocket APIs.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "operation_name",
						Description: "The operation name for the route.",
						Type:        schema.TypeString,
					},
					{
						Name:          "request_models",
						Description:   "The request models for the route. Supported only for WebSocket APIs.",
						Type:          schema.TypeJSON,
						IgnoreInTests: true,
					},
					{
						Name:        "request_parameters",
						Description: "The request parameters for the route. Supported only for WebSocket APIs.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "route_id",
						Description: "The route ID.",
						Type:        schema.TypeString,
					},
					{
						Name:          "route_response_selection_expression",
						Description:   "The route response selection expression for the route. Supported only for WebSocket APIs.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "target",
						Description: "The target for the route.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_apigatewayv2_api_route_responses",
						Description:   "Represents a route response.",
						Resolver:      fetchApigatewayv2ApiRouteResponses,
						Options:       schema.TableCreationOptions{PrimaryKeys: []string{"api_route_cq_id", "route_response_id"}},
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "api_route_cq_id",
								Description: "Unique CloudQuery ID of aws_apigatewayv2_api_routes table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "route_id",
								Description: "Represents the identifier of an route.",
								Type:        schema.TypeString,
								Resolver:    schema.ParentResourceFieldResolver("route_id"),
							},
							{
								Name:        "arn",
								Description: "The Amazon Resource Name (ARN) for the resource.",
								Type:        schema.TypeString,
								Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
									r := resource.Item.(types.RouteResponse)
									route := resource.Parent.Item.(types.Route)
									api := resource.Parent.Parent.Item.(types.Api)
									return []string{apiIDPart, *api.ApiId, apiRouteIDPart, *route.RouteId, "routeresponses", *r.RouteResponseId}, nil
								}),
							},
							{
								Name:        "route_response_key",
								Description: "Represents the route response key of a route response.",
								Type:        schema.TypeString,
							},
							{
								Name:        "model_selection_expression",
								Description: "Represents the model selection expression of a route response. Supported only for WebSocket APIs.",
								Type:        schema.TypeString,
							},
							{
								Name:        "response_models",
								Description: "Represents the response models of a route response.",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "response_parameters",
								Description: "Represents the response parameters of a route response.",
								Type:        schema.TypeJSON,
							},
							{
								Name:        "route_response_id",
								Description: "Represents the identifier of a route response.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_apigatewayv2_api_stages",
				Description: "Represents an API stage.",
				Resolver:    fetchApigatewayv2ApiStages,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"api_cq_id", "stage_name"}},
				Columns: []schema.Column{
					{
						Name:        "api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigatewayv2_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "api_id",
						Description: "The API ID.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Stage)
							p := resource.Parent.Item.(types.Api)
							return []string{apiIDPart, *p.ApiId, "stages", *r.StageName}, nil
						}),
					},
					{
						Name:        "stage_name",
						Description: "The name of the stage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_log_settings_destination_arn",
						Description: "The ARN of the CloudWatch Logs log group to receive access logs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.DestinationArn"),
					},
					{
						Name:        "access_log_settings_format",
						Description: "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least $context.requestId.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.Format"),
					},
					{
						Name:        "api_gateway_managed",
						Description: "Specifies whether a stage is managed by API Gateway. If you created an API using quick create, the $default stage is managed by API Gateway. You can't modify the $default stage.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "auto_deploy",
						Description: "Specifies whether updates to an API automatically trigger a new deployment. The default value is false.",
						Type:        schema.TypeBool,
					},
					{
						Name:          "client_certificate_id",
						Description:   "The identifier of a client certificate for a Stage. Supported only for WebSocket APIs.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "created_date",
						Description: "The timestamp when the stage was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "route_settings_data_trace_enabled",
						Description: "Specifies whether (true) or not (false) data trace logging is enabled for this route. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultRouteSettings.DataTraceEnabled"),
					},
					{
						Name:        "route_settings_detailed_metrics_enabled",
						Description: "Specifies whether detailed metrics are enabled.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DefaultRouteSettings.DetailedMetricsEnabled"),
					},
					{
						Name:        "route_settings_logging_level",
						Description: "Specifies the logging level for this route: INFO, ERROR, or OFF. This property affects the log entries pushed to Amazon CloudWatch Logs. Supported only for WebSocket APIs.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DefaultRouteSettings.LoggingLevel"),
					},
					{
						Name:        "route_settings_throttling_burst_limit",
						Description: "Specifies the throttling burst limit.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DefaultRouteSettings.ThrottlingBurstLimit"),
					},
					{
						Name:        "route_settings_throttling_rate_limit",
						Description: "Specifies the throttling rate limit.",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("DefaultRouteSettings.ThrottlingRateLimit"),
					},
					{
						Name:        "deployment_id",
						Description: "The identifier of the Deployment that the Stage is associated with. Can't be updated if autoDeploy is enabled.",
						Type:        schema.TypeString,
					},
					{
						Name:          "description",
						Description:   "The description of the stage.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "last_deployment_status_message",
						Description: "Describes the status of the last deployment of a stage. Supported only for stages with autoDeploy enabled.",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_updated_date",
						Description: "The timestamp when the stage was last updated.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "route_settings",
						Description: "Route settings for the stage, by routeKey.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "stage_variables",
						Description: "A map that defines the stage variables for a stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "tags",
						Description: "The collection of tags. Each tag element is associated with a given resource.",
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
		response, err := svc.GetApis(ctx, &config, func(o *apigatewayv2.Options) {
			// o.Region = c.Region
		})

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
func fetchApigatewayv2ApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiIntegrations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiIntegrationResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
		return diag.WrapError(err)
	}
	return resource.Set(c.Name, response.Value)
}
func fetchApigatewayv2ApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
func fetchApigatewayv2ApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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
