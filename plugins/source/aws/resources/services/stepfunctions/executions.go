package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func executions() *schema.Table {
	return &schema.Table{
		Name:                "aws_stepfunctions_executions",
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html`,
		Resolver:            fetchStepfunctionsExecutions,
		PreResourceResolver: getExecution,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeExecutionOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExecutionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "state_machine_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
		Relations: []*schema.Table{
			mapRuns(),
		},
	}
}

func fetchStepfunctionsExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	sfnOutput := parent.Item.(*sfn.DescribeStateMachineOutput)
	config := sfn.ListExecutionsInput{
		MaxResults:      1000,
		StateMachineArn: sfnOutput.StateMachineArn,
	}
	paginator := sfn.NewListExecutionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Executions
	}
	return nil
}

func getExecution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	execution := resource.Item.(types.ExecutionListItem)
	svc := meta.(*client.Client).Services().Sfn

	executionResult, err := svc.DescribeExecution(ctx, &sfn.DescribeExecutionInput{
		ExecutionArn: execution.ExecutionArn,
	})
	if err != nil {
		return err
	}
	resource.Item = executionResult
	return nil
}
