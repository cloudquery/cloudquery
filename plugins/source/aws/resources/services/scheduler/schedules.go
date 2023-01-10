package scheduler

import (
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Schedules() *schema.Table {
	return &schema.Table{
		Name:                "aws_scheduler_schedules",
		Description:         `https://docs.aws.amazon.com/scheduler/latest/APIReference/API_GetScheduleOutput.html`,
		Resolver:            fetchSchedulerSchedules,
		PreResourceResolver: getSchedule,
		Transform:           transformers.TransformWithStruct(&scheduler.GetScheduleOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer("scheduler"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveSchedulerScheduleTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
