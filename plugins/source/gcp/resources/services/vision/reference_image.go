package vision

import (
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ReferenceImages() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vision_product_reference_images",
		Description: `https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products.referenceImages`,
		Resolver:    fetchReferenceImages,
		Transform:   client.TransformWithStruct(&pb.ReferenceImage{}, transformers.WithPrimaryKeys("Name")),
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
