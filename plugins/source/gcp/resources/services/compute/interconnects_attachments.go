package compute

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"

	compute "cloud.google.com/go/compute/apiv1"
)

func InterconnectAttachments() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_interconnect_attachments",
		Description: "https://cloud.google.com/compute/docs/reference/rest/v1/interconnectAttachments#InterconnectAttachment",
		Resolver:    fetchInterconnectAttachments,
		Multiplex:   client.ProjectMultiplexEnabledServices("compute.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.InterconnectAttachment{}, transformers.WithPrimaryKeys("SelfLink")),
		Columns: []schema.Column{
			client.ProjectIDColumn(false),
		},
	}
}

func fetchInterconnectAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListInterconnectAttachmentsRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewInterconnectAttachmentsRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.List(ctx, req, c.CallOptions...)
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
