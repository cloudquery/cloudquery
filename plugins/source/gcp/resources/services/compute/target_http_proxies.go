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

func TargetHttpProxies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_target_http_proxies",
		Resolver:  fetchTargetHttpProxies,
		Multiplex: client.ProjectMultiplex("compute.googleapis.com"),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
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
				Name:     "proxy_bind",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ProxyBind"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "url_map",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UrlMap"),
			},
		},
	}
}

func fetchTargetHttpProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListTargetHttpProxiesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewTargetHttpProxiesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.TargetHttpProxies

	}
	return nil
}
