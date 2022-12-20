// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func UrlMaps() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_url_maps",
		Description: `https://cloud.google.com/compute/docs/reference/rest/v1/urlMaps#UrlMap`,
		Resolver:    fetchUrlMaps,
		Multiplex:   client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "default_route_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultRouteAction"),
			},
			{
				Name:     "default_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultService"),
			},
			{
				Name:     "default_url_redirect",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DefaultUrlRedirect"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "header_action",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HeaderAction"),
			},
			{
				Name:     "host_rules",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HostRules"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "path_matchers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PathMatchers"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "self_link",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfLink"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tests",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tests"),
			},
		},
	}
}

func fetchUrlMaps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListUrlMapsRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewUrlMapsRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.UrlMaps

	}
	return nil
}
