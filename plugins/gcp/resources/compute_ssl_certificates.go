package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"google.golang.org/api/compute/v1"
)

func ComputeSslCertificates() *schema.Table {
	return &schema.Table{
		Name:         "gcp_compute_ssl_certificates",
		Resolver:     fetchComputeSslCertificates,
		Multiplex:    client.ProjectMultiplex,
		DeleteFilter: client.DeleteProjectFilter,
		IgnoreError:  client.IgnoreErrorHandler,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "certificate",
				Type: schema.TypeString,
			},
			{
				Name: "creation_timestamp",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "expire_time",
				Type: schema.TypeString,
			},
			{
				Name:     "resource_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveResourceId,
			},
			{
				Name: "kind",
				Type: schema.TypeString,
			},
			{
				Name:     "managed_domain_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Managed.DomainStatus"),
			},
			{
				Name:     "managed_domains",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Managed.Domains"),
			},
			{
				Name:     "managed_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Managed.Status"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "private_key",
				Type: schema.TypeString,
			},
			{
				Name: "region",
				Type: schema.TypeString,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
			},
			{
				Name:     "self_managed_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfManaged.Certificate"),
			},
			{
				Name:     "self_managed_private_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SelfManaged.PrivateKey"),
			},
			{
				Name: "subject_alternative_names",
				Type: schema.TypeStringArray,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeSslCertificates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		call := c.Services.Compute.SslCertificates.AggregatedList(c.ProjectId).Context(ctx)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		var sslCertificate []*compute.SslCertificate
		for _, items := range output.Items {
			sslCertificate = append(sslCertificate, items.SslCertificates...)
		}

		res <- sslCertificate
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
