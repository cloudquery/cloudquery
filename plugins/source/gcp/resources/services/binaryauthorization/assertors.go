package binaryauthorization

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
)

func Assertors() *schema.Table {
	return &schema.Table{
		Name:        "gcp_binaryauthorization_assertors",
		Description: `https://cloud.google.com/binary-authorization/docs/reference/rest/v1/projects.attestors#Attestor`,
		Resolver:    fetchAssertors,
		Multiplex:   client.ProjectMultiplexEnabledServices("binaryauthorization.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Attestor{}, client.Options()...),
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

func fetchAssertors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListAttestorsRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := binaryauthorization.NewBinauthzManagementClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListAttestors(ctx, req, c.CallOptions...)
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
