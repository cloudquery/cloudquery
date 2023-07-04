package billing

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	billing "cloud.google.com/go/billing/apiv1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_billing_services",
		Description: `https://cloud.google.com/billing/docs/reference/rest/v1/services/list#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplex,
		Transform:   client.TransformWithStruct(&pb.Service{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListServicesRequest{}
	gcpClient, err := billing.NewCloudCatalogClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
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
