package compute

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeTargetSslProxies() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_target_ssl_proxies",
		Description:  "Represents a Target SSL Proxy resource",
		Resolver:     fetchComputeTargetSslProxies,
		Multiplex:    client.ProjectMultiplex,
		IgnoreError:  client.IgnoreErrorHandler,
		DeleteFilter: client.DeleteProjectFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
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
				Resolver:    client.ISODateResolver("CreationTimestamp"),
			},
			{
				Name:        "description",
				Description: "An optional description of this resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The unique identifier for the resource This identifier is defined by the server",
				Type:        schema.TypeString,
				Resolver:    client.ResolveResourceId,
			},
			{
				Name:        "kind",
				Description: "Type of the resource Always compute#targetSslProxy for target SSL proxies",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Name of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "proxy_header",
				Description: "Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1 The default is NONE",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "service",
				Description: "URL to the BackendService resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssl_certificates",
				Description: "URLs to SslCertificate resources that are used to authenticate connections to Backends At least one SSL certificate must be specified Currently, you may specify up to 15 SSL certificates",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "ssl_policy",
				Description: "URL of SslPolicy resource that will be associated with the TargetSslProxy resource If not set, the TargetSslProxy resource will not have any SSL policy configured",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeTargetSslProxies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.TargetSslProxies.List(c.ProjectId).PageToken(nextPageToken)
		list, err := c.RetryingDo(ctx, call)
		if err != nil {
			return err
		}
		output := list.(*compute.TargetSslProxyList)

		res <- output.Items

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
