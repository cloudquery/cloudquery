package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_repositories",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories#Repository`,
		Resolver:    fetchRepositories,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Repository{}, client.Options()...),
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
		Relations: []*schema.Table{
			DockerImages(),
			Files(),
			Packages(),
		},
	}
}
