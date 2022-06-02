package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

const restApiIDPart = "/restapis"

func ApigatewayRestApis() *schema.Table {
	return &schema.Table{
		Name:          "aws_apigateway_rest_apis",
		Description:   "Represents a REST API.",
		Resolver:      fetchApigatewayRestApis,
		Multiplex:     client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
		IgnoreInTests: true,
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
					return []string{restApiIDPart, *resource.Item.(types.RestApi).Id}, nil
				}),
			},
			{
				Name:        "api_key_source",
				Description: "The source of the API key for metering requests according to a usage plan. Valid values are:",
				Type:        schema.TypeString,
			},
			{
				Name:        "binary_media_types",
				Description: "The list of binary media types supported by the RestApi. By default, the RestApi supports only UTF-8-encoded text payloads.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The API's description.",
				Type:        schema.TypeString,
			},
			{
				Name:        "disable_execute_api_endpoint",
				Description: "Specifies whether clients can invoke your API by using the default execute-api endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "endpoint_configuration_types",
				Description: "A list of endpoint types of an API (RestApi) or its custom domain name (DomainName). For an edge-optimized API and its custom domain name, the endpoint type is \"EDGE\". For a regional API and its custom domain name, the endpoint type is REGIONAL. For a private API, the endpoint type is PRIVATE.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:        "endpoint_configuration_vpc_endpoint_ids",
				Description: "A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes. It is only supported for PRIVATE endpoint type.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
			},
			{
				Name:        "id",
				Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
			},
			{
				Name:        "minimum_compression_size",
				Description: "A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API. When compression is enabled, compression or decompression is not applied on the payload if the payload size is smaller than this value. Setting it to zero allows compression for any payload size.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "name",
				Description: "The API's name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags. Each tag element is associated with a given resource.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "version",
				Description: "A version identifier for the API.",
				Type:        schema.TypeString,
			},
			{
				Name:        "warnings",
				Description: "The warning messages reported when failonwarnings is turned on during API import.",
				Type:        schema.TypeStringArray,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_apigateway_rest_api_authorizers",
				Description:   "Represents an authorization layer for methods.",
				Resolver:      fetchApigatewayRestApiAuthorizers,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Authorizer)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "authorizers", *r.Id}, nil
						}),
					},
					{
						Name:        "auth_type",
						Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact.",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_credentials",
						Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_result_ttl_in_seconds",
						Description: "The TTL in seconds of cached authorizer results. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway will cache authorizer responses. If this field is not set, the default value is 300. The maximum value is 3600, or 1 hour.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "authorizer_uri",
						Description: "Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or REQUEST authorizers, this must be a well-formed Lambda function URI, for example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations. In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api}, where {region} is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial /. For Lambda functions, this is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the authorizer resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "identity_source",
						Description: "The identity source for which authorization is requested.",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_validation_expression",
						Description: "A validation expression for the incoming identity token. For TOKEN authorizers, this value is a regular expression. For COGNITO_USER_POOLS authorizers, API Gateway will match the aud field of the incoming token from the client against the specified regular expression. It will invoke the authorizer's Lambda function when there is a match. Otherwise, it will return a 401 Unauthorized response without calling the Lambda function. The validation expression does not apply to the REQUEST authorizer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "[Required] The name of the authorizer.",
						Type:        schema.TypeString,
					},
					{
						Name:        "provider_arns",
						Description: "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer. Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}. For a TOKEN or REQUEST authorizer, this is not defined.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ProviderARNs"),
					},
					{
						Name:        "type",
						Description: "The authorizer type. Valid values are TOKEN for a Lambda function using a single authorization token submitted in a custom header, REQUEST for a Lambda function using incoming request parameters, and COGNITO_USER_POOLS for using an Amazon Cognito user pool.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_deployments",
				Description:   "An immutable representation of a RestApi resource that can be called by users using Stages.",
				Resolver:      fetchApigatewayRestApiDeployments,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Deployment)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "deployments", *r.Id}, nil
						}),
					},
					{
						Name:        "api_summary",
						Description: "A summary of the RestApi at the date and time that the deployment resource was created.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "created_date",
						Description: "The date and time that the deployment resource was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description for the deployment resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the deployment resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_documentation_parts",
				Description:   "A documentation part for a targeted API entity.",
				Resolver:      fetchApigatewayRestApiDocumentationParts,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.DocumentationPart)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "documentation/parts", *r.Id}, nil
						}),
					},
					{
						Name:        "id",
						Description: "The DocumentationPart identifier, generated by API Gateway when the DocumentationPart is created.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "location_type",
						Description: "[Required] The type of API entity to which the documentation content applies. Valid values are API, AUTHORIZER, MODEL, RESOURCE, METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. Content inheritance does not apply to any entity of the API, AUTHORIZER, METHOD, MODEL, REQUEST_BODY, or RESOURCE type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Type"),
					},
					{
						Name:        "location_method",
						Description: "The HTTP verb of a method. It is a valid field for the API entity types of METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is * for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other location attributes, the child entity's method attribute must match that of the parent entity exactly.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Method"),
					},
					{
						Name:        "location_name",
						Description: "The name of the targeted API entity. It is a valid and required field for the API entity types of AUTHORIZER, MODEL, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY and RESPONSE_HEADER. It is an invalid field for any other entity type.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Name"),
					},
					{
						Name:        "location_path",
						Description: "The URL path of the target. It is a valid field for the API entity types of RESOURCE, METHOD, PATH_PARAMETER, QUERY_PARAMETER, REQUEST_HEADER, REQUEST_BODY, RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is / for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other location attributes, the child entity's path attribute must match that of the parent entity as a prefix.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Path"),
					},
					{
						Name:        "location_status_code",
						Description: "The HTTP status code of a response. It is a valid field for the API entity types of RESPONSE, RESPONSE_HEADER, and RESPONSE_BODY. The default value is * for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other location attributes, the child entity's statusCode attribute must match that of the parent entity exactly.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.StatusCode"),
					},
					{
						Name:        "properties",
						Description: "A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., \"{ \\\"description\\\": \\\"The API does ...\\\" }\". Only OpenAPI-compliant documentation-related fields from the properties map are exported and, hence, published as part of the API entity definitions, while the original documentation parts are exported in a OpenAPI extension of x-amazon-apigateway-documentation.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_documentation_versions",
				Description:   "A snapshot of the documentation of an API.",
				Resolver:      fetchApigatewayRestApiDocumentationVersions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.DocumentationVersion)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "documentation/versions", *r.Version}, nil
						}),
					},
					{
						Name:        "created_date",
						Description: "The date when the API documentation snapshot is created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description of the API documentation snapshot.",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The version identifier of the API documentation snapshot.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_gateway_responses",
				Description:   "A gateway response of a given response type and status code, with optional response parameters and mapping templates.",
				Resolver:      fetchApigatewayRestApiGatewayResponses,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.GatewayResponse)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "gatewayresponses", string(r.ResponseType)}, nil
						}),
					},
					{
						Name:        "default_response",
						Description: "A Boolean flag to indicate whether this GatewayResponse is the default gateway response (true) or not (false). A default gateway response is one generated by API Gateway without any customization by an API developer.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "response_parameters",
						Description: "Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "response_templates",
						Description: "Response templates of the GatewayResponse as a string-to-string map of key-value pairs.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "response_type",
						Description: "The response type of the associated GatewayResponse.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_code",
						Description: "The HTTP status code for this GatewayResponse.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_models",
				Description:   "Represents the data structure of a method's request or response payload.",
				Resolver:      fetchApigatewayRestApiModels,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Model)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "models", *r.Name}, nil
						}),
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayRestAPIModelModelTemplate,
					},
					{
						Name:        "content_type",
						Description: "The content-type for the model.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the model.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the model resource.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "name",
						Description: "The name of the model. Must be an alphanumeric string.",
						Type:        schema.TypeString,
					},
					{
						Name:        "schema",
						Description: "The schema for the model. For application/json models, this should be JSON schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Do not include \"\\*/\" characters in the description of any properties because such \"\\*/\" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_request_validators",
				Description:   "A set of validation rules for incoming Method requests.",
				Resolver:      fetchApigatewayRestApiRequestValidators,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.RequestValidator)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "requestvalidators", *r.Id}, nil
						}),
					},
					{
						Name:        "id",
						Description: "The identifier of this RequestValidator.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "name",
						Description: "The name of this RequestValidator",
						Type:        schema.TypeString,
					},
					{
						Name:        "validate_request_body",
						Description: "A Boolean flag to indicate whether to validate a request body according to the configured Model schema.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "validate_request_parameters",
						Description: "A Boolean flag to indicate whether to validate request parameters (true) or not (false).",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_resources",
				Description:   "Represents an API resource.",
				Resolver:      fetchApigatewayRestApiResources,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Resource)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "resources", *r.Id}, nil
						}),
					},
					{
						Name:        "id",
						Description: "The resource's identifier.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Id"),
					},
					{
						Name:        "parent_id",
						Description: "The parent resource's identifier.",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The full path for this resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "path_part",
						Description: "The last path segment for this resource.",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_methods",
						Description: "Gets an API resource's method of a given HTTP verb. The resource methods are a map of methods indexed by methods' HTTP verbs enabled on the resource. This method map is included in the 200 OK response of the GET /restapis/{restapi_id}/resources/{resource_id} or GET /restapis/{restapi_id}/resources/{resource_id}?embed=methods request. Example: Get the GET method of an API resource",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:          "aws_apigateway_rest_api_stages",
				Description:   "Represents a unique identifier for a version of a deployed RestApi that is callable by users.",
				Resolver:      fetchApigatewayRestApiStages,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier. This identifier is unique across all of your APIs in API Gateway.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARNWithRegion(client.ApigatewayService, func(resource *schema.Resource) ([]string, error) {
							r := resource.Item.(types.Stage)
							p := resource.Parent.Item.(types.RestApi)
							return []string{restApiIDPart, *p.Id, "stages", *r.StageName}, nil
						}),
					},
					{
						Name:        "access_log_settings_destination_arn",
						Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with amazon-apigateway-.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.DestinationArn"),
					},
					{
						Name:        "access_log_settings_format",
						Description: "A single line format of the access logs of data, as specified by selected $context variables (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least $context.requestId.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("AccessLogSettings.Format"),
					},
					{
						Name:        "cache_cluster_enabled",
						Description: "Specifies whether a cache cluster is enabled for the stage.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "cache_cluster_size",
						Description: "The size of the cache cluster for the stage, if enabled.",
						Type:        schema.TypeString,
					},
					{
						Name:        "cache_cluster_status",
						Description: "The status of the cache cluster for the stage, if enabled.",
						Type:        schema.TypeString,
					},
					{
						Name:        "canary_settings_deployment_id",
						Description: "The ID of the canary deployment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CanarySettings.DeploymentId"),
					},
					{
						Name:        "canary_settings_percent_traffic",
						Description: "The percent (0-100) of traffic diverted to a canary deployment.",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("CanarySettings.PercentTraffic"),
					},
					{
						Name:        "canary_settings_stage_variable_overrides",
						Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("CanarySettings.StageVariableOverrides"),
					},
					{
						Name:        "canary_settings_use_stage_cache",
						Description: "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CanarySettings.UseStageCache"),
					},
					{
						Name:        "client_certificate_id",
						Description: "The identifier of a client certificate for an API stage.",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_date",
						Description: "The timestamp when the stage was created.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "deployment_id",
						Description: "The identifier of the Deployment that the stage points to.",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The stage's description.",
						Type:        schema.TypeString,
					},
					{
						Name:        "documentation_version",
						Description: "The version of the associated API documentation.",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_updated_date",
						Description: "The timestamp when the stage last updated.",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "method_settings",
						Description: "A map that defines the method settings for a Stage resource. Keys (designated as /{method_setting_key below) are method paths defined as {resource_path}/{http_method} for an individual method override, or /\\*/\\* for overriding all methods in the stage.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "stage_name",
						Description: "The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The collection of tags. Each tag element is associated with a given resource.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "tracing_enabled",
						Description: "Specifies whether active tracing with X-ray is enabled for the Stage.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "variables",
						Description: "A map that defines the stage variables for a Stage resource. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "web_acl_arn",
						Description: "The ARN of the WebAcl associated with the Stage.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchApigatewayRestApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config apigateway.GetRestApisInput
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	for {
		response, err := svc.GetRestApis(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiAuthorizers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetAuthorizersInput{RestApiId: r.Id}
	for {
		response, err := svc.GetAuthorizers(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiDeployments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDeploymentsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDeployments(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiDocumentationParts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationPartsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationParts(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiDocumentationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetDocumentationVersionsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetDocumentationVersions(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiGatewayResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetGatewayResponsesInput{RestApiId: r.Id}
	for {
		response, err := svc.GetGatewayResponses(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiModels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetModelsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetModels(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return diag.WrapError(err)
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

	response, err := svc.GetModelTemplate(ctx, &config, func(options *apigateway.Options) {
		options.Region = cl.Region
	})
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
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, response.Value))
}
func fetchApigatewayRestApiRequestValidators(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetRequestValidatorsInput{RestApiId: r.Id}
	for {
		response, err := svc.GetRequestValidators(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetResourcesInput{RestApiId: r.Id}
	for {
		response, err := svc.GetResources(ctx, &config, func(options *apigateway.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.Position) == "" {
			break
		}
		config.Position = response.Position
	}
	return nil
}
func fetchApigatewayRestApiStages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.RestApi)
	c := meta.(*client.Client)
	svc := c.Services().Apigateway
	config := apigateway.GetStagesInput{RestApiId: r.Id}

	response, err := svc.GetStages(ctx, &config, func(options *apigateway.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	res <- response.Item

	return nil
}
