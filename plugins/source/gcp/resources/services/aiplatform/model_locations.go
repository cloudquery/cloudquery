package aiplatform

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	pb "google.golang.org/genproto/googleapis/cloud/location"

	"google.golang.org/api/option"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
)

func ModelLocations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_model_locations",
		Description: `https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location`,
		Resolver:    fetchModelLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Location{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			Models(),
		},
	}
}

func fetchModelLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListLocationsRequest{
		Name: "projects/" + c.ProjectId,
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewModelClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListLocations(ctx, req, c.CallOptions...)
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
