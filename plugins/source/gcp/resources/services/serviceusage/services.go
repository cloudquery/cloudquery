// Code generated by codegen; DO NOT EDIT.

package serviceusage

import (
	"context"
	"google.golang.org/api/iterator"

	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/serviceusage/apiv1"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "gcp_serviceusage_services",
		Description: `https://cloud.google.com/service-usage/docs/reference/rest/v1/services#Service`,
		Resolver:    fetchServices,
		Multiplex:   client.ProjectMultiplex,
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
				Name:     "parent",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Parent"),
			},
			{
				Name:     "config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Config"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: client.ResolveProtoEnum("State"),
			},
		},
	}
}

func fetchServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.ListServicesRequest{
		Parent:   "projects/" + c.ProjectId,
		PageSize: 200,
		Filter:   "state:ENABLED",
	}
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
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
