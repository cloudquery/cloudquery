package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_repositories",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository`,
		Resolver:    fetchRepositories,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Repository{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			client.ProjectIDColumn(true),
		},
		Relations: []*schema.Table{
			DockerImages(),
			Files(),
			Packages(),
		},
	}
}
