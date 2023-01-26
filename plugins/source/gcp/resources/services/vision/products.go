package vision

import (
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Products() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vision_products",
		Description: `https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products`,
		Resolver:    fetchProducts,
		// The list of locations was copied by me from the APIs error message. Yuck.
		Multiplex: client.ProjectLocationMultiplexEnabledServices("vision.googleapis.com",
			[]string{
				"us-west1", "us-east1", "asia-east1", "europe-west1",
			}),
		Transform: transformers.TransformWithStruct(&pb.Product{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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

		Relations: []*schema.Table{
			ReferenceImages(),
		},
	}
}
