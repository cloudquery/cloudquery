package lambda

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
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
				Name: "name",
				Type: arrow.BinaryTypes.String,
				Resolver: func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
					return r.Set(c.Name, string(r.Item.(types.Runtime)))
				},
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLambdaRuntimes(_ context.Context, _ schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	res <- types.Runtime("").Values()
	return nil
}
