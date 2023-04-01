package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func functionVersions() *schema.Table {
	tableName := "aws_lambda_function_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_FunctionConfiguration.html`,
		Resolver:    fetchLambdaFunctionVersions,
		Transform:   transformers.TransformWithStruct(&types.FunctionConfiguration{}),
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

func fetchLambdaFunctionVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	p := parent.Item.(*lambda.GetFunctionOutput)
	if p.Configuration == nil {
		return nil
	}

	svc := meta.(*client.Client).Services().Lambda
	config := lambda.ListVersionsByFunctionInput{
		FunctionName: p.Configuration.FunctionName,
	}

	for {
		output, err := svc.ListVersionsByFunction(ctx, &config)
		if err != nil {
			if meta.(*client.Client).IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.Versions
		if output.NextMarker == nil {
			break
		}
		config.Marker = output.NextMarker
	}
	return nil
}
