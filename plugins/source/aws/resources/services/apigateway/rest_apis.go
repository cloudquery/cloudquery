package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource rest_apis --config rest_apis.hcl --output .
func RestApis() *schema.Table {
	return &schema.Table{
		Name:         "aws_apigateway_rest_apis",
		Description:  "Represents a REST API",
		Resolver:     fetchApigatewayRestApis,
		Multiplex:    client.ServiceAccountRegionMultiplexer("apigateway"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource",
				Type:        schema.TypeString,
				Resolver:    resolveApigatewayRestAPIArn,
			},
			{
				Name:        "api_key_source",
				Description: "The source of the API key for metering requests according to a usage plan",
				Type:        schema.TypeString,
			},
			{
				Name:        "binary_media_types",
				Description: "The list of binary media types supported by the RestApi",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "created_date",
				Description: "The timestamp when the API was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "description",
				Description: "The API's description",
				Type:        schema.TypeString,
			},
			{
				Name:        "disable_execute_api_endpoint",
				Description: "Specifies whether clients can invoke your API by using the default execute-api endpoint",
				Type:        schema.TypeBool,
			},
			{
				Name:        "endpoint_configuration_types",
				Description: "A list of endpoint types of an API (RestApi) or its custom domain name (DomainName)",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.Types"),
			},
			{
				Name:        "endpoint_configuration_vpc_endpoint_ids",
				Description: "A list of VpcEndpointIds of an API (RestApi) against which to create Route53 ALIASes",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("EndpointConfiguration.VpcEndpointIds"),
			},
			{
				Name:        "id",
				Description: "The API's identifier",
				Type:        schema.TypeString,
			},
			{
				Name:        "minimum_compression_size",
				Description: "A nullable integer that is used to enable compression (with non-negative between 0 and 10485760 (10M) bytes, inclusive) or disable compression (with a null value) on an API",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "The API's name",
				Type:        schema.TypeString,
			},
			{
				Name:        "policy",
				Description: "A stringified JSON policy document that applies to this RestApi regardless of the caller and Method configuration",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The collection of tags",
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
				Name:        "aws_apigateway_rest_api_authorizers",
				Description: "Represents an authorization layer for methods",
				Resolver:    fetchApigatewayRestApiAuthorizers,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIAuthorizerArn,
					},
					{
						Name:        "auth_type",
						Description: "Optional customer-defined field, used in OpenAPI imports and exports without functional impact",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_credentials",
						Description: "Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "authorizer_result_ttl_in_seconds",
						Description: "The TTL in seconds of cached authorizer results",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "authorizer_uri",
						Description: "Specifies the authorizer's Uniform Resource Identifier (URI)",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the authorizer resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_source",
						Description: "The identity source for which authorization is requested",
						Type:        schema.TypeString,
					},
					{
						Name:        "identity_validation_expression",
						Description: "A validation expression for the incoming identity token",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the authorizer",
						Type:        schema.TypeString,
					},
					{
						Name:        "provider_arns",
						Description: "A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ProviderARNs"),
					},
					{
						Name:        "type",
						Description: "The authorizer type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_deployments",
				Description: "An immutable representation of a RestApi resource that can be called by users using Stages",
				Resolver:    fetchApigatewayRestApiDeployments,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIDeploymentArn,
					},
					{
						Name:        "api_summary",
						Description: "A summary of the RestApi at the date and time that the deployment resource was created",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "created_date",
						Description: "The date and time that the deployment resource was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description for the deployment resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the deployment resource",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_documentation_parts",
				Description: "A documentation part for a targeted API entity",
				Resolver:    fetchApigatewayRestApiDocumentationParts,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIDocumentationPartArn,
					},
					{
						Name:        "id",
						Description: "The DocumentationPart identifier, generated by API Gateway when the DocumentationPart is created",
						Type:        schema.TypeString,
					},
					{
						Name:        "location_type",
						Description: "The type of API entity to which the documentation content applies",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Type"),
					},
					{
						Name:        "location_method",
						Description: "The HTTP verb of a method",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Method"),
					},
					{
						Name:        "location_name",
						Description: "The name of the targeted API entity",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Name"),
					},
					{
						Name:        "location_path",
						Description: "The URL path of the target",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.Path"),
					},
					{
						Name:        "location_status_code",
						Description: "The HTTP status code of a response",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.StatusCode"),
					},
					{
						Name:        "properties",
						Description: "A content map of API-specific key-value pairs describing the targeted API entity",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_documentation_versions",
				Description: "A snapshot of the documentation of an API",
				Resolver:    fetchApigatewayRestApiDocumentationVersions,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIDocumentationVersionArn,
					},
					{
						Name:        "created_date",
						Description: "The date when the API documentation snapshot is created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "description",
						Description: "The description of the API documentation snapshot",
						Type:        schema.TypeString,
					},
					{
						Name:        "version",
						Description: "The version identifier of the API documentation snapshot",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_gateway_responses",
				Description: "A gateway response of a given response type and status code, with optional response parameters and mapping templates",
				Resolver:    fetchApigatewayRestApiGatewayResponses,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIGatewayResponseArn,
					},
					{
						Name:        "default_response",
						Description: "A Boolean flag to indicate whether this GatewayResponse is the default gateway response (true) or not (false)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "response_parameters",
						Description: "Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "response_templates",
						Description: "Response templates of the GatewayResponse as a string-to-string map of key-value pairs",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "response_type",
						Description: "The response type of the associated GatewayResponse",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_code",
						Description: "The HTTP status code for this GatewayResponse",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_models",
				Description: "Represents the data structure of a method's request or response payload",
				Resolver:    fetchApigatewayRestApiModels,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIModelArn,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: resolveApigatewayRestAPIModelModelTemplate,
					},
					{
						Name:        "content_type",
						Description: "The content-type for the model",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The description of the model",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "The identifier for the model resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the model",
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
				Name:        "aws_apigateway_rest_api_request_validators",
				Description: "A set of validation rules for incoming Method requests",
				Resolver:    fetchApigatewayRestApiRequestValidators,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIRequestValidatorArn,
					},
					{
						Name:        "id",
						Description: "The identifier of this RequestValidator",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of this RequestValidator",
						Type:        schema.TypeString,
					},
					{
						Name:        "validate_request_body",
						Description: "A Boolean flag to indicate whether to validate a request body according to the configured Model schema",
						Type:        schema.TypeBool,
					},
					{
						Name:        "validate_request_parameters",
						Description: "A Boolean flag to indicate whether to validate request parameters (true) or not (false)",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_resources",
				Description: "Represents an API resource",
				Resolver:    fetchApigatewayRestApiResources,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIResourceArn,
					},
					{
						Name:        "id",
						Description: "The resource's identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "parent_id",
						Description: "The parent resource's identifier",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The full path for this resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "path_part",
						Description: "The last path segment for this resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_methods",
						Description: "Gets an API resource's method of a given HTTP verb",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "aws_apigateway_rest_api_stages",
				Description: "Represents a unique identifier for a version of a deployed RestApi that is callable by users",
				Resolver:    fetchApigatewayRestApiStages,
				Columns: []schema.Column{
					{
						Name:        "rest_api_cq_id",
						Description: "Unique CloudQuery ID of aws_apigateway_rest_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rest_api_id",
						Description: "The API's identifier",
						Type:        schema.TypeString,
						Resolver:    schema.ParentPathResolver("Id"),
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource",
						Type:        schema.TypeString,
						Resolver:    resolveApigatewayRestAPIStageArn,
					},
					{
						Name:        "access_log_settings_destination_arn",
						Description: "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs",
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
						Name:        "cache_cluster_enabled",
						Description: "Specifies whether a cache cluster is enabled for the stage",
						Type:        schema.TypeBool,
					},
					{
						Name:        "cache_cluster_size",
						Description: "The size of the cache cluster for the stage, if enabled",
						Type:        schema.TypeString,
					},
					{
						Name:        "cache_cluster_status",
						Description: "The status of the cache cluster for the stage, if enabled",
						Type:        schema.TypeString,
					},
					{
						Name:        "canary_settings_deployment_id",
						Description: "The ID of the canary deployment",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CanarySettings.DeploymentId"),
					},
					{
						Name:        "canary_settings_percent_traffic",
						Description: "The percent (0-100) of traffic diverted to a canary deployment",
						Type:        schema.TypeFloat,
						Resolver:    schema.PathResolver("CanarySettings.PercentTraffic"),
					},
					{
						Name:        "canary_settings_stage_variable_overrides",
						Description: "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("CanarySettings.StageVariableOverrides"),
					},
					{
						Name:        "canary_settings_use_stage_cache",
						Description: "A Boolean flag to indicate whether the canary deployment uses the stage cache or not",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("CanarySettings.UseStageCache"),
					},
					{
						Name:        "client_certificate_id",
						Description: "The identifier of a client certificate for an API stage",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_date",
						Description: "The timestamp when the stage was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "deployment_id",
						Description: "The identifier of the Deployment that the stage points to",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "The stage's description",
						Type:        schema.TypeString,
					},
					{
						Name:        "documentation_version",
						Description: "The version of the associated API documentation",
						Type:        schema.TypeString,
					},
					{
						Name:        "last_updated_date",
						Description: "The timestamp when the stage last updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "method_settings",
						Description: "A map that defines the method settings for a Stage resource",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "stage_name",
						Description: "The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The collection of tags",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "tracing_enabled",
						Description: "Specifies whether active tracing with X-ray is enabled for the Stage",
						Type:        schema.TypeBool,
					},
					{
						Name:        "variables",
						Description: "A map that defines the stage variables for a Stage resource",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "web_acl_arn",
						Description: "The ARN of the WebAcl associated with the Stage",
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
	for p := apigateway.NewGetRestApisPaginator(svc, &config); p.HasMorePages(); {
		response, err := p.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
	}
	return nil
}
func resolveApigatewayRestAPIArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	rapi := resource.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
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
func resolveApigatewayRestAPIAuthorizerArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	auth := resource.Item.(types.Authorizer)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "authorizers", *auth.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
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
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, arn))
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
func resolveApigatewayRestAPIDocumentationPartArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	d := resource.Item.(types.DocumentationPart)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "documentation/parts", *d.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
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
func resolveApigatewayRestAPIDocumentationVersionArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	v := resource.Item.(types.DocumentationVersion)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "documentation/versions", *v.Version)
	return diag.WrapError(resource.Set(c.Name, arn))
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
func resolveApigatewayRestAPIGatewayResponseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.GatewayResponse)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "gatewayresponses", string(r.ResponseType))
	return diag.WrapError(resource.Set(c.Name, arn))
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
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, arn))
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
		response, err := svc.GetRequestValidators(ctx, &config)
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
func resolveApigatewayRestAPIRequestValidatorArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	r := resource.Item.(types.RequestValidator)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "requestvalidators", *r.Id)
	return diag.WrapError(resource.Set(c.Name, arn))
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
			return diag.WrapError(err)
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
	return diag.WrapError(resource.Set(c.Name, arn))
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
		return diag.WrapError(err)
	}
	res <- response.Item

	return nil
}
func resolveApigatewayRestAPIStageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	s := resource.Item.(types.Stage)
	rapi := resource.Parent.Item.(types.RestApi)
	arn := cl.RegionGlobalARN(client.ApigatewayService, restApiIDPart, *rapi.Id, "stages", *s.StageName)
	return diag.WrapError(resource.Set(c.Name, arn))
}
