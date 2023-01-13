package beyondcorp

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	appgateways "cloud.google.com/go/beyondcorp/appgateways/apiv1"
)

func AppGateways() *schema.Table {
	return &schema.Table{
		Name:        "gcp_beyondcorp_app_gateways",
		Description: `https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appGateways#AppGateway`,
		Resolver:    fetchAppGateways,
		Multiplex:   client.ProjectMultiplexEnabledServices("beyondcorp.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.AppGateway{}, client.Options()...),
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

func fetchAppGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListAppGatewaysRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := appgateways.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAppGateways(ctx, req, c.CallOptions...)
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
