package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FunctionVersions() *schema.Table {
	tableName := "aws_lambda_function_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html`,
		Resolver:    fetchLambdaFunctionVersions,
		Transform:   client.TransformWithStruct(&types.FunctionConfiguration{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
