package lambda

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func functionConcurrencyConfigs() *schema.Table {
	tableName := "aws_lambda_function_concurrency_configs"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_ProvisionedConcurrencyConfigListItem.html`,
		Resolver:    fetchLambdaFunctionConcurrencyConfigs,
		Transform:   transformers.TransformWithStruct(&types.ProvisionedConcurrencyConfigListItem{}),
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

func fetchLambdaFunctionConcurrencyConfigs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceLambda).Lambda
	config := lambda.ListProvisionedConcurrencyConfigsInput{
		FunctionName: p.Configuration.FunctionName,
	}
	paginator := lambda.NewListProvisionedConcurrencyConfigsPaginator(svc, &config)
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
		res <- page.ProvisionedConcurrencyConfigs
	}
	return nil
}
