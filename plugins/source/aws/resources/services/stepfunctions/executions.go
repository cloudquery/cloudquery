package stepfunctions

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func executions() *schema.Table {
	tableName := "aws_stepfunctions_executions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html`,
		Resolver:            fetchStepfunctionsExecutions,
		PreResourceResolver: getExecution,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeExecutionOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ExecutionArn"),
				PrimaryKey: true,
			},
			{
				Name:     "state_machine_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
		Relations: []*schema.Table{
			mapRuns(),
		},
	}
}

func fetchStepfunctionsExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Sfn
	sfnOutput := parent.Item.(*sfn.DescribeStateMachineOutput)
	config := sfn.ListExecutionsInput{
		MaxResults:      1000,
		StateMachineArn: sfnOutput.StateMachineArn,
	}
	paginator := sfn.NewListExecutionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *sfn.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Executions
	}
	return nil
}

func getExecution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	execution := resource.Item.(types.ExecutionListItem)

	cl := meta.(*client.Client)
	svc := cl.Services().Sfn

	executionResult, err := svc.DescribeExecution(ctx, &sfn.DescribeExecutionInput{
		ExecutionArn: execution.ExecutionArn,
	}, func(o *sfn.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = executionResult
	return nil
}
