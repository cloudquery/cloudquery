package apigateway

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	apigateway "cloud.google.com/go/apigateway/apiv1"
)

func Gateways() *schema.Table {
	return &schema.Table{
		Name:        "gcp_apigateway_gateways",
		Description: `https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations.gateways#Gateway`,
		Resolver:    fetchGateways,
		Multiplex:   client.ProjectMultiplexEnabledServices("apigateway.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Gateway{}, transformers.WithPrimaryKeys("Name")),
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

func fetchGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListGatewaysRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := apigateway.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListGateways(ctx, req, c.CallOptions...)
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
