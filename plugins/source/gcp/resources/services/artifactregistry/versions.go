package artifactregistry

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_versions",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.versions#Version`,
		Resolver:    fetchVersions,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Version{}, client.Options()...),
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

func fetchVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListVersionsRequest{
		Parent: parent.Item.(*pb.Package).Name,
	}
	gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListVersions(ctx, req, c.CallOptions...)
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
