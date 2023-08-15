package scheduler

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ScheduleGroups() *schema.Table {
	tableName := "aws_scheduler_schedule_groups"
	return &schema.Table{
		Name:        tableName,
		Title:       "Amazon EventBridge Scheduler Schedule Groups",
		Description: `https://docs.aws.amazon.com/scheduler/latest/APIReference/API_ScheduleGroupSummary.html`,
		Resolver:    fetchSchedulerScheduleGroups,
		Transform:   transformers.TransformWithStruct(&types.ScheduleGroupSummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "scheduler"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveSchedulerScheduleTags(),
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchSchedulerScheduleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := scheduler.ListScheduleGroupsInput{
		MaxResults: aws.Int32(100),
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Scheduler
	paginator := scheduler.NewListScheduleGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *scheduler.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.ScheduleGroups
	}
	return nil
}
