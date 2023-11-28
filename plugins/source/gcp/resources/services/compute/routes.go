package compute

import (
	"context"
	"errors"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/iterator"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Routes() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_routes",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/routes/list#response-body`,
		Resolver:    fetchRoutes,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&computepb.Route{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &computepb.ListRoutesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewRoutesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}

	it := gcpClient.List(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return err
		}
		res <- resp
	}

	return nil
}
