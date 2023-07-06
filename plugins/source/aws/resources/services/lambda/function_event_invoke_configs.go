package lambda

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func functionEventInvokeConfigs() *schema.Table {
	tableName := "aws_lambda_function_event_invoke_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionEventInvokeConfig.html`,
		Resolver:    fetchLambdaFunctionEventInvokeConfigs,
		Transform:   transformers.TransformWithStruct(&types.FunctionEventInvokeConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "function_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}
	paginator := lambda.NewListFunctionEventInvokeConfigsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *lambda.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- page.FunctionEventInvokeConfigs
	}
	return nil
}
