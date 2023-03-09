package scheduler

import (
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ScheduleGroups() *schema.Table {
	tableName := "aws_scheduler_schedule_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/scheduler/latest/APIReference/API_ScheduleGroupSummary.html`,
		Resolver:    fetchSchedulerScheduleGroups,
		Transform:   transformers.TransformWithStruct(&types.ScheduleGroupSummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "scheduler"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
