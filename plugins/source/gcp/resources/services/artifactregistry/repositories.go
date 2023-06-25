package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_repositories",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository`,
		Resolver:    fetchRepositories,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Repository{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			DockerImages(),
			Files(),
			Packages(),
		},
	}
}
