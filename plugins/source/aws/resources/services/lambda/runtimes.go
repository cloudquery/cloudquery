package lambda

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func Runtimes() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_runtimes",
		Description: "https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html",
		Resolver:    fetchLambdaRuntimes,
		Columns: schema.ColumnList{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	for _, runtime := range types.Runtime("").Values() {
		res <- struct{ Name string }{Name: string(runtime)}
	}
	return nil
}
