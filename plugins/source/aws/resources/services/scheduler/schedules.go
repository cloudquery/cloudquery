package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/thoas/go-funk"
)

func Schedules() *schema.Table {
	tableName := "aws_scheduler_schedules"
	return &schema.Table{
		Name:                tableName,
		Title:               "Amazon EventBridge Scheduler Schedules",
		Description:         `https://docs.aws.amazon.com/scheduler/latest/APIReference/API_GetScheduleOutput.html`,
		Resolver:            fetchSchedulerSchedules,
		PreResourceResolver: getSchedule,
		Transform:           transformers.TransformWithStruct(&scheduler.GetScheduleOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "scheduler"),
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

func fetchSchedulerSchedules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := scheduler.ListSchedulesInput{
		MaxResults: aws.Int32(100),
	}
	c := meta.(*client.Client)
	svc := c.Services().Scheduler
	paginator := scheduler.NewListSchedulesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Schedules
	}
	return nil
}

func getSchedule(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Scheduler
	scheduleSummary := resource.Item.(types.ScheduleSummary)

	describeTaskDefinitionOutput, err := svc.GetSchedule(ctx, &scheduler.GetScheduleInput{
		Name:      scheduleSummary.Name,
		GroupName: scheduleSummary.GroupName,
	})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput
	return nil
}

func resolveSchedulerScheduleTags() schema.ColumnResolver {
	return func(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		arnStr := funk.Get(r.Item, "Arn", funk.WithAllowZero()).(*string)
		cl := meta.(*client.Client)
		svc := cl.Services().Scheduler
		params := scheduler.ListTagsForResourceInput{ResourceArn: arnStr}

		output, err := svc.ListTagsForResource(ctx, &params)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
