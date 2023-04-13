package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	for _, runtime := range types.Runtime("").Values() {
		res <- &models.RuntimeWrapper{Name: string(runtime)}
	}
	return nil
}
