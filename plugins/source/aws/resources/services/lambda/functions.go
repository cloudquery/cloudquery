package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:                 "aws_lambda_functions",
		Description:          `https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html`,
		Resolver:             fetchLambdaFunctions,
		PreResourceResolver:  getFunction,
		PostResourceResolver: resolvePolicyCodeSigningConfig,
		Transform:            transformers.TransformWithStruct(&lambda.GetFunctionOutput{}),
		Multiplex:            client.ServiceAccountRegionMultiplexer("lambda"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.FunctionArn"),
			},
			{
				Name: "policy_revision_id",
				Type: schema.TypeString,
			},
			{
				Name: "policy_document",
				Type: schema.TypeJSON,
			},
			{
				Name: "code_signing_config",
				Type: schema.TypeJSON,
			},
			{
				Name:     "code_repository_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.RepositoryType"),
			},
		},

		Relations: []*schema.Table{
			FunctionEventInvokeConfigs(),
			FunctionAliases(),
			FunctionVersions(),
			FunctionConcurrencyConfigs(),
			FunctionEventSourceMappings(),
		},
	}
}
