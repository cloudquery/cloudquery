package artifactregistry

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/artifactregistry/v1"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_locations",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
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
			Repositories(),
		},
	}
}
