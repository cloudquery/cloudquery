package artifactregistry

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/artifactregistry/v1"
)

func Locations() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_locations",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/Shared.Types/ListLocationsResponse#Location`,
		Resolver:    fetchLocations,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Location{}, append(client.Options(), transformers.WithPrimaryKeys("Name"))...),
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
			Repositories(),
		},
	}
}
