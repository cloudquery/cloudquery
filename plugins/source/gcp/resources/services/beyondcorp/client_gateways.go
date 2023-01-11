package beyondcorp

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	clientgateways "cloud.google.com/go/beyondcorp/clientgateways/apiv1"
)

func ClientGateways() *schema.Table {
	return &schema.Table{
		Name:        "gcp_beyondcorp_client_gateways",
		Description: `https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.clientGateways#ClientGateway`,
		Resolver:    fetchClientGateways,
		Multiplex:   client.ProjectMultiplexEnabledServices("beyondcorp.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.ClientGateway{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchClientGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListClientGatewaysRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := clientgateways.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListClientGateways(ctx, req, c.CallOptions...)
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
