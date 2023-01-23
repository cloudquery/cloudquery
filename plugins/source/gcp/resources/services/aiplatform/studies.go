package aiplatform

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	"google.golang.org/api/option"

	"google.golang.org/genproto/googleapis/cloud/location"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
)

func Studies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_aiplatform_studies",
		Description: `https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.studies#Study`,
		Resolver:    fetchStudies,
		Multiplex:   client.ProjectMultiplexEnabledServices("aiplatform.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Study{}, client.Options()...),
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

func fetchStudies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListStudiesRequest{
		Parent: parent.Item.(*location.Location).Name,
	}
	if filterStudiesLocation(parent) {
		return nil
	}

	clientOptions := c.ClientOptions
	clientOptions = append([]option.ClientOption{option.WithEndpoint(parent.Item.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
	gcpClient, err := aiplatform.NewVizierClient(ctx, clientOptions...)

	if err != nil {
		return err
	}
	it := gcpClient.ListStudies(ctx, req, c.CallOptions...)
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
