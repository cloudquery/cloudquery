package stepfunctions

import (
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Executions() *schema.Table {
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
		},
		Relations: []*schema.Table{
			mapRuns(),
		},
	}
}
