package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Functions() *schema.Table {
	tableName := "aws_lambda_functions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_GetFunction.html`,
		Resolver:            fetchLambdaFunctions,
		PreResourceResolver: getFunction,
		Transform:           transformers.TransformWithStruct(&lambda.GetFunctionOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Configuration.FunctionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name: "policy_revision_id",
				Type: schema.TypeString,
				// resolved in resolveResourcePolicy
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: resolveResourcePolicy,
			},
			{
				Name:     "code_signing_config",
				Type:     schema.TypeJSON,
				Resolver: resolveCodeSigningConfig,
			},
			{
				Name:     "code_repository_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Code.RepositoryType"),
			},
			{
				Name: "update_runtime_on",
				Type: schema.TypeString,
				// resolved in resolveRuntimeManagementConfig
			},
			{
				Name:     "runtime_version_arn",
				Type:     schema.TypeString,
				Resolver: resolveRuntimeManagementConfig,
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
