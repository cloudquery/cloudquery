package cloudscheduler

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"

	scheduler "cloud.google.com/go/scheduler/apiv1"
	"cloud.google.com/go/scheduler/apiv1/schedulerpb"

	locationspb "google.golang.org/api/cloudscheduler/v1"
)

func fetchJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	p := parent.Item.(*locationspb.Location)
	gcpClient, err := scheduler.NewCloudSchedulerClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.ListJobs(ctx, &schedulerpb.ListJobsRequest{Parent: p.Name})
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		res <- resp
	}
	return nil
}
