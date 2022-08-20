package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
)

func ComputeTargetHTTPSProxies() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_target_https_proxies",
		Description: "Represents a Target HTTPS Proxy resource",
		Resolver:    fetchComputeTargetHttpsProxies,
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
				Name:        "authorization_policy",
				Description: "A URL referring to a networksecurityAuthorizationPolicy resource that describes how the proxy should authorize inbound traffic If left blank, access will not be restricted by an authorization policy Refer to the AuthorizationPolicy resource for additional details authorizationPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED Note: This field currently has no impact",
				Type:        schema.TypeString,
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
				Description: "Unique Id of the ssl proxy",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of resource Always compute#targetHttpsProxy for target HTTPS proxies",
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
				Name:        "quic_override",
				Description: "Specifies the QUIC override policy for this TargetHttpsProxy resource This setting determines whether the load balancer attempts to negotiate QUIC with clients You can specify NONE, ENABLE, or DISABLE - When quic-override is set to NONE, Google manages whether QUIC is used - When quic-override is set to ENABLE, the load balancer uses QUIC when possible - When quic-override is set to DISABLE, the load balancer doesn't use QUIC - If the quic-override flag is not specified, NONE is implied",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional TargetHttpsProxy resides This field is not applicable to global TargetHttpsProxies",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "server_tls_policy",
				Description: "A URL referring to a networksecurityServerTlsPolicy resource that describes how the proxy should authenticate inbound traffic serverTlsPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED If left blank, communications are not encrypted Note: This field currently has no impact",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssl_certificates",
				Description: "URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "ssl_policy",
				Description: "URL of SslPolicy resource that will be associated with the TargetHttpsProxy resource If not set, the TargetHttpsProxy resource has no SSL policy configured",
				Type:        schema.TypeString,
			},
			{
				Name:        "url_map",
				Description: "A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService For example, the following are all valid URLs for specifying a URL map: - https://wwwgoogleapis",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeTargetHttpsProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.TargetHttpsProxies.List(c.ProjectId).PageToken(nextPageToken).Do()
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
