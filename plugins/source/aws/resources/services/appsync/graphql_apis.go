package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource graphql_apis --config gen.hcl --output .
func GraphqlApis() *schema.Table {
	return &schema.Table{
		Name:         "aws_appsync_graphql_apis",
		Description:  "Describes a GraphQL API",
		Resolver:     fetchAppsyncGraphqlApis,
		Multiplex:    client.ServiceAccountRegionMultiplexer("appsync"),
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
				Name:        "id",
				Description: "The API ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ApiId"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN)",
				Type:        schema.TypeString,
			},
			{
				Name:        "authentication_type",
				Description: "The authentication type",
				Type:        schema.TypeString,
			},
			{
				Name:        "lambda_authorizer_config_authorizer_uri",
				Description: "The Amazon Resource Name (ARN) of the Lambda function to be called for authorization",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaAuthorizerConfig.AuthorizerUri"),
			},
			{
				Name:        "lambda_authorizer_config_authorizer_result_ttl_in_seconds",
				Description: "The number of seconds a response should be cached for",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("LambdaAuthorizerConfig.AuthorizerResultTtlInSeconds"),
			},
			{
				Name:        "lambda_authorizer_config_identity_validation_expression",
				Description: "A regular expression for validation of tokens before the Lambda function is called",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LambdaAuthorizerConfig.IdentityValidationExpression"),
			},
			{
				Name:        "log_config_cloud_watch_logs_role_arn",
				Description: "The service role that AppSync assumes to publish to CloudWatch logs in your account",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.CloudWatchLogsRoleArn"),
			},
			{
				Name:        "log_config_field_log_level",
				Description: "The field logging level",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LogConfig.FieldLogLevel"),
			},
			{
				Name:        "log_config_exclude_verbose_content",
				Description: "Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging level",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LogConfig.ExcludeVerboseContent"),
			},
			{
				Name:        "name",
				Description: "The API name",
				Type:        schema.TypeString,
			},
			{
				Name:        "open_id_connect_config_issuer",
				Description: "The issuer for the OIDC configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OpenIDConnectConfig.Issuer"),
			},
			{
				Name:        "open_id_connect_config_auth_ttl",
				Description: "The number of milliseconds that a token is valid after being authenticated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OpenIDConnectConfig.AuthTTL"),
			},
			{
				Name:        "open_id_connect_config_client_id",
				Description: "The client identifier of the relying party at the OpenID identity provider",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("OpenIDConnectConfig.ClientId"),
			},
			{
				Name:        "open_id_connect_config_iat_ttl",
				Description: "The number of milliseconds that a token is valid after it's issued to a user",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("OpenIDConnectConfig.IatTTL"),
			},
			{
				Name:        "tags",
				Description: "The tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "uris",
				Description: "The URIs",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "user_pool_config_aws_region",
				Description: "The Amazon Web Services Region in which the user pool was created",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserPoolConfig.AwsRegion"),
			},
			{
				Name:        "user_pool_config_default_action",
				Description: "The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user pool configuration",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserPoolConfig.DefaultAction"),
			},
			{
				Name:        "user_pool_config_user_pool_id",
				Description: "The user pool ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserPoolConfig.UserPoolId"),
			},
			{
				Name:        "user_pool_config_app_id_client_regex",
				Description: "A regular expression for validating the incoming Amazon Cognito user pool app client ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("UserPoolConfig.AppIdClientRegex"),
			},
			{
				Name:        "waf_web_acl_arn",
				Description: "The ARN of the WAF access control list (ACL) associated with this GraphqlApi, if one exists",
				Type:        schema.TypeString,
			},
			{
				Name:        "xray_enabled",
				Description: "A flag indicating whether to use X-Ray tracing for this GraphqlApi",
				Type:        schema.TypeBool,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_appsync_graphql_api_additional_authentication_providers",
				Description: "Describes an additional authentication provider",
				Resolver:    schema.PathTableResolver("AdditionalAuthenticationProviders"),
				Columns: []schema.Column{
					{
						Name:        "graphql_api_cq_id",
						Description: "Unique CloudQuery ID of aws_appsync_graphql_apis table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "authentication_type",
						Description: "The authentication type: API key, Identity and Access Management (IAM), OpenID Connect (OIDC), Amazon Cognito user pools, or Lambda",
						Type:        schema.TypeString,
					},
					{
						Name:        "lambda_authorizer_config_authorizer_uri",
						Description: "The Amazon Resource Name (ARN) of the Lambda function to be called for authorization",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LambdaAuthorizerConfig.AuthorizerUri"),
					},
					{
						Name:        "lambda_authorizer_config_authorizer_result_ttl_in_seconds",
						Description: "The number of seconds a response should be cached for",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("LambdaAuthorizerConfig.AuthorizerResultTtlInSeconds"),
					},
					{
						Name:        "lambda_authorizer_config_identity_validation_expression",
						Description: "A regular expression for validation of tokens before the Lambda function is called",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LambdaAuthorizerConfig.IdentityValidationExpression"),
					},
					{
						Name:        "open_id_connect_config_issuer",
						Description: "The issuer for the OIDC configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenIDConnectConfig.Issuer"),
					},
					{
						Name:        "open_id_connect_config_auth_ttl",
						Description: "The number of milliseconds that a token is valid after being authenticated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("OpenIDConnectConfig.AuthTTL"),
					},
					{
						Name:        "open_id_connect_config_client_id",
						Description: "The client identifier of the relying party at the OpenID identity provider",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("OpenIDConnectConfig.ClientId"),
					},
					{
						Name:        "open_id_connect_config_iat_ttl",
						Description: "The number of milliseconds that a token is valid after it's issued to a user",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("OpenIDConnectConfig.IatTTL"),
					},
					{
						Name:        "user_pool_config_aws_region",
						Description: "The Amazon Web Services Region in which the user pool was created",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UserPoolConfig.AwsRegion"),
					},
					{
						Name:        "user_pool_config_user_pool_id",
						Description: "The user pool ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UserPoolConfig.UserPoolId"),
					},
					{
						Name:        "user_pool_config_app_id_client_regex",
						Description: "A regular expression for validating the incoming Amazon Cognito user pool app client ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("UserPoolConfig.AppIdClientRegex"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAppsyncGraphqlApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config appsync.ListGraphqlApisInput
	c := meta.(*client.Client)
	svc := c.Services().AppSync
	for {
		output, err := svc.ListGraphqlApis(ctx, &config, func(options *appsync.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.GraphqlApis
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
