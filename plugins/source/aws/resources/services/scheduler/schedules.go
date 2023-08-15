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

func fetchSchedulerSchedules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := scheduler.ListSchedulesInput{
		MaxResults: aws.Int32(100),
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Scheduler
	paginator := scheduler.NewListSchedulesPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *scheduler.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Schedules
	}
	return nil
}

func getSchedule(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Scheduler
	scheduleSummary := resource.Item.(types.ScheduleSummary)

	describeTaskDefinitionOutput, err := svc.GetSchedule(ctx, &scheduler.GetScheduleInput{
		Name:      scheduleSummary.Name,
		GroupName: scheduleSummary.GroupName,
	}, func(o *scheduler.Options) {
		o.Region = cl.Region
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

		output, err := svc.ListTagsForResource(ctx, &params, func(o *scheduler.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		return r.Set(c.Name, client.TagsToMap(output.Tags))
	}
}
