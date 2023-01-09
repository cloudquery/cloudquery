package containeranalysis

import (
	"context"

	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
)

func Occurrences() *schema.Table {
	return &schema.Table{
		Name:        "gcp_containeranalysis_occurrences",
		Description: `https://cloud.google.com/container-analysis/docs/reference/rest/v1beta1/projects.occurrences#Occurrence`,
		Resolver:    fetchOccurrences,
		Multiplex:   client.ProjectMultiplexEnabledServices("containeranalysis.googleapis.com"),
		Transform:   transformers.TransformWithStruct(&pb.Occurrence{}, client.Options()...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
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

func fetchOccurrences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListOccurrencesRequest{
		Parent: "projects/" + c.ProjectId,
	}
	gcpClient, err := containeranalysis.NewGrafeasV1Beta1Client(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListOccurrences(ctx, req, c.CallOptions...)
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
