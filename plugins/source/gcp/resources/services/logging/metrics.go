package logging

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	logging "cloud.google.com/go/logging/apiv2"
)

func Metrics() *schema.Table {
	return &schema.Table{
		Name:        "gcp_logging_metrics",
		Description: `https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics#LogMetric`,
		Resolver:    fetchMetrics,
		Multiplex:   client.ProjectMultiplexEnabledServices("logging.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.LogMetric{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

func fetchMetrics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListLogMetricsRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := logging.NewMetricsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListLogMetrics(ctx, req, c.CallOptions...)
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
