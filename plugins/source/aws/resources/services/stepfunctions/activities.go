package stepfunctions

import (
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Activities() *schema.Table {
	return &schema.Table{
		Name:        "aws_stepfunctions_activities",
		Description: `https://docs.aws.amazon.com/step-functions/latest/apireference/API_DescribeStateMachine.html`,
		Resolver:    fetchStepfunctionsActivities,
		Transform:   transformers.TransformWithStruct(&types.ActivityListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("states"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ActivityArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{},
	}
}
