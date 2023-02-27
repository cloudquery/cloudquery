package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Runtimes() *schema.Table {
	return &schema.Table{
		Name:     "aws_lambda_runtimes",
		Resolver: fetchLambdaRuntimes,
		Transform: transformers.TransformWithStruct(&models.RuntimeWrapper{},
			transformers.WithPrimaryKeys("Name"),
		),
	}
}
