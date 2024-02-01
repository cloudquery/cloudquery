package lambda

import (
	"context"

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
		Transform: transformers.TransformWithStruct(&types.FunctionEventInvokeConfig{},
			transformers.WithPrimaryKeyComponents("FunctionArn"), // FunctionArn here can also be a version
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
		},
	}
}
func fetchLambdaFunctionEventInvokeConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	config := lambda.ListFunctionEventInvokeConfigsInput{FunctionName: p.Configuration.FunctionArn}
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
