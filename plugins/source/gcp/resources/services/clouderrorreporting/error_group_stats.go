package clouderrorreporting

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
)

func ErrorGroupStats() *schema.Table {
	return &schema.Table{
		Name:        "gcp_clouderrorreporting_error_group_stats",
		Description: `https://cloud.google.com/error-reporting/reference/rest/v1beta1/projects.groupStats/list#ErrorGroupStats`,
		Resolver:    fetchErrorGroupStats,
		Multiplex:   client.ProjectMultiplexEnabledServices("clouderrorreporting.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ErrorGroupStats{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
		},
		Relations: []*schema.Table{
			ErrorEvents(),
		},
	}
}

func fetchErrorGroupStats(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListGroupStatsRequest{
		ProjectName: "projects/" + c.ProjectId,
	}
	gcpClient, err := errorreporting.NewErrorStatsClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListGroupStats(ctx, req, c.CallOptions...)
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
