package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func mapRunExecutions() *schema.Table {
	return &schema.Table{
		Name:                "aws_stepfunctions_map_run_executions",
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeExecution.html`,
		Resolver:            fetchStepfunctionsMapRunExecutions,
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
				Name:     "map_run_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "state_machine_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("state_machine_arn"),
			},
		},
	}
}

func fetchStepfunctionsMapRunExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListExecutionsInput{
		MaxResults: 1000,
		MapRunArn:  parent.Item.(*sfn.DescribeMapRunOutput).MapRunArn,
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
