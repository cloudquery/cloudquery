package stepfunctions

import (
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func mapRuns() *schema.Table {
	return &schema.Table{
		Name:                "aws_stepfunctions_executions_map_runs",
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
		},
		Relations: []*schema.Table{},
	}
}
