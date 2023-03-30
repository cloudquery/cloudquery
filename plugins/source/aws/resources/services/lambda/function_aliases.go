package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FunctionAliases() *schema.Table {
	tableName := "aws_lambda_function_aliases"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/lambda/latest/dg/API_AliasConfiguration.html`,
		Resolver:            fetchLambdaFunctionAliases,
		PreResourceResolver: getFunctionAliasURLConfig,
		Transform:           client.TransformWithStruct(&models.AliasWrapper{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
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
