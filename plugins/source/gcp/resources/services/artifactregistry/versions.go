package artifactregistry

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
)

func Versions() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_versions",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.packages.versions#Version`,
		Resolver:    fetchVersions,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Version{}, transformers.WithPrimaryKeys("Name")),
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
