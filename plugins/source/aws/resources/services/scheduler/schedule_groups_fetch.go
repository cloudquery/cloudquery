package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSchedulerScheduleGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := scheduler.ListScheduleGroupsInput{
		MaxResults: aws.Int32(100),
	}
	c := meta.(*client.Client)
	svc := c.Services().Scheduler
	paginator := scheduler.NewListScheduleGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.ScheduleGroups
	}
	return nil
}
