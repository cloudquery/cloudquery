// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"

	"google.golang.org/api/compute/v1"
)

func ComputeTargetHttpProxies() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_target_http_proxies",
		Resolver:  fetchComputeTargetHttpProxies,
		Multiplex: client.ProjectMultiplex,
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
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "url_map",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UrlMap"),
			},
		},
	}
}

func fetchComputeTargetHttpProxies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.TargetHttpProxies.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		var allItems []*compute.TargetHttpProxy
		for _, items := range output.Items {
			allItems = append(allItems, items.TargetHttpProxies...)
		}
		res <- allItems

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
