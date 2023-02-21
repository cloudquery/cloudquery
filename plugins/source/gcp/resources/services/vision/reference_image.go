package vision

import (
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func ReferenceImages() *schema.Table {
	return &schema.Table{
		Name:        "gcp_vision_product_reference_images",
		Description: `https://cloud.google.com/vision/docs/reference/rest/v1/projects.locations.products.referenceImages`,
		Resolver:    fetchReferenceImages,
		Transform:   transformers.TransformWithStruct(&pb.ReferenceImage{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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
