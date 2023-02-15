package stepfunctions

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func mapRuns() *schema.Table {
	return &schema.Table{
		Name:                "aws_stepfunctions_map_runs",
		Description:         `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeMapRun.html`,
		Resolver:            fetchStepfunctionsMapRuns,
		PreResourceResolver: getMapRun,
		Transform:           transformers.TransformWithStruct(&sfn.DescribeMapRunOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MapRunArn"),
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
			mapRunExecutions(),
		},
	}
}

func fetchStepfunctionsMapRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.ListMapRunsInput{
		MaxResults:   1000,
		ExecutionArn: parent.Item.(*sfn.DescribeExecutionOutput).ExecutionArn,
	}
	paginator := sfn.NewListMapRunsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.MapRuns
	}
	return nil
}

func getMapRun(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Sfn
	config := sfn.DescribeMapRunInput{
		MapRunArn: resource.Item.(types.MapRunListItem).MapRunArn,
	}
	output, err := svc.DescribeMapRun(ctx, &config)
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}
