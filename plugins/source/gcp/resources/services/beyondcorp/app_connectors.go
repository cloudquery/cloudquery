package beyondcorp

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/beyondcorp/appconnectors/apiv1/appconnectorspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	appconnectors "cloud.google.com/go/beyondcorp/appconnectors/apiv1"
)

func AppConnectors() *schema.Table {
	return &schema.Table{
		Name:        "gcp_beyondcorp_app_connectors",
		Description: `https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnectors#AppConnector`,
		Resolver:    fetchAppConnectors,
		Multiplex:   client.ProjectMultiplexEnabledServices("beyondcorp.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.AppConnector{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchAppConnectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListAppConnectorsRequest{
		Parent: "projects/" + c.ProjectId + "/locations/-",
	}
	gcpClient, err := appconnectors.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAppConnectors(ctx, req, c.CallOptions...)
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
