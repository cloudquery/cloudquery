package vision

import (
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
		Transform: client.TransformWithStruct(&pb.Product{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			ReferenceImages(),
		},
	}
}
