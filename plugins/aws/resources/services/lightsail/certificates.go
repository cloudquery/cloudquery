package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource certificates --config gen.hcl --output .
func Certificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_certificates",
		Description:  "Describes the full details of an Amazon Lightsail SSL/TLS certificate",
		Resolver:     fetchLightsailCertificates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the certificate was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "domain_name",
				Description: "The domain name of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "eligible_to_renew",
				Description: "The renewal eligibility of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "in_use_resource_count",
				Description: "The number of Lightsail resources that the certificate is attached to",
				Type:        schema.TypeInt,
			},
			{
				Name:          "issued_at",
				Description:   "The timestamp when the certificate was issued",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "issuer_ca",
				Description: "The certificate authority that issued the certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("IssuerCA"),
			},
			{
				Name:        "key_algorithm",
				Description: "The algorithm used to generate the key pair (the public and private key) of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the certificate (eg, my-certificate)",
				Type:        schema.TypeString,
			},
			{
				Name:          "not_after",
				Description:   "The timestamp when the certificate expires",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "not_before",
				Description:   "The timestamp when the certificate is first valid",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "renewal_summary_status",
				Description: "The renewal status of the certificate",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RenewalSummary.RenewalStatus"),
			},
			{
				Name:          "renewal_summary_reason",
				Description:   "The reason for the renewal status of the certificate",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("RenewalSummary.RenewalStatusReason"),
				IgnoreInTests: true,
			},
			{
				Name:          "renewal_summary_updated_at",
				Description:   "The timestamp when the certificate was last updated",
				Type:          schema.TypeTimestamp,
				Resolver:      schema.PathResolver("RenewalSummary.UpdatedAt"),
				IgnoreInTests: true,
			},
			{
				Name:        "request_failure_reason",
				Description: "The validation failure reason, if any, of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:          "revocation_reason",
				Description:   "The reason the certificate was revoked",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "revoked_at",
				Description:   "The timestamp when the certificate was revoked",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:          "serial_number",
				Description:   "The serial number of the certificate",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "The validation status of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "subject_alternative_names",
				Description: "An array of strings that specify the alternate domains (eg, example2com) and subdomains (eg, blogexamplecom) of the certificate",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    resolveCertificatesTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lightsail_certificate_domain_validation_records",
				Description:   "Describes the domain validation records of an Amazon Lightsail SSL/TLS certificate",
				Resolver:      fetchLightsailCertificateDomainValidationRecords,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "certificate_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_certificates table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "domain_name",
						Description: "The domain name of the certificate validation record",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the record",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Name"),
					},
					{
						Name:        "type",
						Description: "The DNS record type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Type"),
					},
					{
						Name:        "value",
						Description: "The value for the DNS record",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Value"),
					},
				},
			},
			{
				Name:          "aws_lightsail_certificate_renewal_summary_domain_validation_records",
				Description:   "Describes the domain validation records of an Amazon Lightsail SSL/TLS certificate",
				Resolver:      fetchLightsailCertificateRenewalSummaryDomainValidationRecords,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "certificate_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_certificates table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "domain_name",
						Description: "The domain name of the certificate validation record",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "The name of the record",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Name"),
					},
					{
						Name:        "type",
						Description: "The DNS record type",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Type"),
					},
					{
						Name:        "value",
						Description: "The value for the DNS record",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ResourceRecord.Value"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := lightsail.GetCertificatesInput{
		IncludeCertificateDetails: true,
	}
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	response, err := svc.GetCertificates(ctx, &input, func(options *lightsail.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	for _, cer := range response.Certificates {
		res <- cer.CertificateDetail
	}
	return nil
}
func resolveCertificatesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*types.Certificate)
	tags := make(map[string]string)
	client.TagsIntoMap(r.Tags, tags)
	return diag.WrapError(resource.Set(c.Name, tags))
}
func fetchLightsailCertificateDomainValidationRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*types.Certificate)
	res <- r.DomainValidationRecords
	return nil
}
func fetchLightsailCertificateRenewalSummaryDomainValidationRecords(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(*types.Certificate)
	if r.RenewalSummary == nil {
		return nil
	}
	res <- r.RenewalSummary.DomainValidationRecords
	return nil
}
