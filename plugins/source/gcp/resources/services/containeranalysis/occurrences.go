package containeranalysis

import (
	"context"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/api/iterator"
)

func Occurrences() *schema.Table {
	return &schema.Table{
		Name:        "gcp_containeranalysis_occurrences",
		Description: `https://cloud.google.com/container-analysis/docs/reference/rest/v1beta1/projects.occurrences#Occurrence`,
		Resolver:    fetchOccurrences,
		Multiplex:   client.ProjectMultiplexEnabledServices("containeranalysis.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.Occurrence{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "details",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveDetails,
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

// details column resolver
func resolveDetails(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, col schema.Column) error {
	occurrence := r.Item.(*pb.Occurrence)
	return r.Set(col.Name, occurrence.GetDetails())
}
