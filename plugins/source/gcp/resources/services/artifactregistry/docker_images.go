// Code generated by codegen; DO NOT EDIT.

package artifactregistry

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/artifactregistry/apiv1"
)

func DockerImages() *schema.Table {
	return &schema.Table{
		Name:        "gcp_artifactregistry_docker_images",
		Description: `https://cloud.google.com/artifact-registry/docs/reference/rest/v1/projects.locations.repositories.dockerImages#DockerImage`,
		Resolver:    fetchDockerImages,
		Multiplex:   client.ProjectMultiplexEnabledServices("artifactregistry.googleapis.com"),
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
			{
				Name:     "uri",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Uri"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "image_size_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ImageSizeBytes"),
			},
			{
				Name:     "upload_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UploadTime"),
			},
			{
				Name:     "media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MediaType"),
			},
			{
				Name:     "build_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("BuildTime"),
			},
		},
	}
}

func fetchDockerImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListDockerImagesRequest{
		Parent: parent.Item.(*pb.Repository).Name,
	}
	gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListDockerImages(ctx, req, c.CallOptions...)
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
