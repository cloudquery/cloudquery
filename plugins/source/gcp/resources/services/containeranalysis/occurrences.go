// Code generated by codegen; DO NOT EDIT.

package containeranalysis

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
)

func Occurrences() *schema.Table {
	return &schema.Table{
		Name:      "gcp_containeranalysis_occurrences",
		Resolver:  fetchOccurrences,
		Multiplex: client.ProjectMultiplex,
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
			{
				Name:     "resource",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Resource"),
			},
			{
				Name:     "note_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NoteName"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("Kind"),
			},
			{
				Name:     "remediation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Remediation"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreateTime"),
			},
			{
				Name:     "update_time",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("UpdateTime"),
			},
		},
	}
}

func fetchOccurrences(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.ListOccurrencesRequest{
		Parent: "projects/" + c.ProjectId,
	}
	it := c.Services.ContaineranalysisGrafeasV1Beta1Client.ListOccurrences(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return errors.WithStack(err)
		}

		res <- resp

	}
	return nil
}
