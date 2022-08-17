package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeTargetHTTPProxies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_target_http_proxies",
		Description: "Represents a Target HTTP Proxy resource",
		Resolver:    fetchComputeTargetHttpProxies,
		Multiplex:   client.ProjectMultiplex,

		Options: schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "project_id",
				Description: "GCP Project Id of the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveProject,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.DateResolver("CreationTimestamp"),
			},
			{
				Name:        "description",
				Description: "An optional description of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "fingerprint",
				Description: "Fingerprint of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of resource Always compute#targetHttpProxy for target HTTP proxies",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "proxy_bind",
				Description: "This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED  When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy) The Envoy proxy listens for inbound requests and handles requests when it receives them  The default is false",
				Type:        schema.TypeBool,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "url_map",
				Description: "A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService.",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeTargetHttpProxies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.TargetHttpProxies.List(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
		}

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
