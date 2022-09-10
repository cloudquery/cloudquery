package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GraphqlApis() *schema.Table {
	return &schema.Table{
		Name:        "aws_appsync_graphql_apis",
		Description: "Describes a GraphQL API",
		Resolver:    fetchAppsyncGraphqlApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("appsync"),
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
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN)",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "authentication_type",
				Description: "The authentication type",
				Type:        schema.TypeString,
			},
			{
				Name:     "lambda_authorizer_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LambdaAuthorizerConfig"),
			},
			{
				Name:     "log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogConfig"),
			},
			{
				Name:        "name",
				Description: "The API name",
				Type:        schema.TypeString,
			},
			{
				Name:     "open_id_connect_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OpenIDConnectConfig"),
			},
			{
				Name:        "open_id_connect_config_iat_ttl",
				Description: "The number of milliseconds that a token is valid after it's issued to a user",
				Type:        schema.TypeInt,
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
				Name:     "user_pool_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolConfig"),
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
			{
				Name:        "additional_authentication_providers",
				Description: "Describes an additional authentication provider",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("AdditionalAuthenticationProviders"),
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
			return err
		}
		res <- output.GraphqlApis
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
