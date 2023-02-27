package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func FunctionAliases() *schema.Table {
	return &schema.Table{
		Name:                "aws_lambda_function_aliases",
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_AliasConfiguration.html`,
		Resolver:            fetchLambdaFunctionAliases,
		PreResourceResolver: getFunctionAliasURLConfig,
		Transform:           transformers.TransformWithStruct(&models.AliasWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer("lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AliasArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
