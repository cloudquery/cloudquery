package compute

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/pkg/errors"
	"google.golang.org/api/compute/v1"
)

func ComputeSslCertificates() *schema.Table {
	return &schema.Table{
		Name:        "gcp_compute_ssl_certificates",
		Description: "Represents an SSL Certificate resource.",
		Resolver:    fetchComputeSslCertificates,
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
				Name:        "certificate",
				Description: "A value read into memory from a certificate file The certificate file must be in PEM format The certificate chain must be no greater than 5 certs long The chain must include at least one intermediate cert",
				Type:        schema.TypeString,
			},
			{
				Name:        "creation_timestamp",
				Description: "Creation timestamp in RFC3339 text format",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "An optional description of this resource Provide this property when you create the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "expire_time",
				Description: "Expire time of the certificate",
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
				Description: "Type of the resource Always compute#sslCertificate for SSL certificates",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_domain_status",
				Description: "[Output only] Detailed statuses of the domains specified for managed certificate resource",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Managed.DomainStatus"),
			},
			{
				Name:        "managed_domains",
				Description: "The domains for which a managed SSL certificate will be generated Each Google-managed SSL certificate supports up to the maximum number of domains per Google-managed SSL certificate (/load-balancing/docs/quotas#ssl_certificates)",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Managed.Domains"),
			},
			{
				Name:        "managed_status",
				Description: "[Output only] Status of the managed certificate resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Managed.Status"),
			},
			{
				Name:        "name",
				Description: "Name of the resource Provided by the client when the resource is created",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_key",
				Description: "A value read into memory from a write-only private key file The private key file must be in PEM format For security, only insert requests include this field",
				Type:        schema.TypeString,
			},
			{
				Name:        "region",
				Description: "URL of the region where the regional SSL Certificate resides This field is not applicable to global SSL Certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_link",
				Description: "[Output only] Server-defined URL for the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "self_managed_certificate",
				Description: "A local certificate file The certificate must be in PEM format The certificate chain must be no greater than 5 certs long The chain must include at least one intermediate cert",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfManaged.Certificate"),
			},
			{
				Name:        "self_managed_private_key",
				Description: "A write-only private key in PEM format Only insert requests will include this field",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SelfManaged.PrivateKey"),
			},
			{
				Name:          "subject_alternative_names",
				Description:   "Domains associated with the certificate via Subject Alternative Name",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "type",
				Description: "Specifies the type of SSL certificate, either \"SELF_MANAGED\" or \"MANAGED\" If not specified, the certificate is self-managed and the fields certificate and private_key are used",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchComputeSslCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextPageToken := ""
	for {
		output, err := c.Services.Compute.SslCertificates.AggregatedList(c.ProjectId).PageToken(nextPageToken).Do()
		if err != nil {
			return errors.WithStack(err)
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
