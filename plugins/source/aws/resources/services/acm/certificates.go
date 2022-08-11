package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func AcmCertificates() *schema.Table {
	return &schema.Table{
		Name:          "aws_acm_certificates",
		Description:   "Contains metadata about an ACM certificate",
		Resolver:      fetchAcmCertificates,
		Multiplex:     client.ServiceAccountRegionMultiplexer("acm"),
		IgnoreError:   client.IgnoreCommonErrors,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		IgnoreInTests: true,
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
				Resolver:    schema.PathResolver("CertificateArn"),
			},
			{
				Name:        "certificate_authority_arn",
				Description: "The Amazon Resource Name (ARN) of the ACM PCA private certificate authority (CA) that issued the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The time at which the certificate was requested.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "domain_name",
				Description: "The fully qualified domain name for the certificate, such as www.example.com or example.com.",
				Type:        schema.TypeString,
			},
			{
				Name:        "domain_validation_options",
				Description: "Contains information about the initial validation of each domain name that occurs as a result of the RequestCertificate request.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DomainValidationOptions"),
			},
			{
				Name:        "extended_key_usages",
				Description: "Contains a list of Extended Key Usage X.509 v3 extension objects.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ExtendedKeyUsages"),
			},
			{
				Name:        "failure_reason",
				Description: "The reason the certificate request failed",
				Type:        schema.TypeString,
			},
			{
				Name:        "imported_at",
				Description: "The date and time at which the certificate was imported",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "in_use_by",
				Description: "A list of ARNs for the Amazon Web Services resources that are using the certificate",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "issued_at",
				Description: "The time at which the certificate was issued",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "issuer",
				Description: "The name of the certificate authority that issued and signed the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_algorithm",
				Description: "The algorithm that was used to generate the public-private key pair.",
				Type:        schema.TypeString,
			},
			{
				Name:        "key_usages",
				Description: "A list of Key Usage X.509 v3 extension objects. Each object is a string value that identifies the purpose of the public key contained in the certificate.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("KeyUsages.Name"),
			},
			{
				Name:        "not_after",
				Description: "The time after which the certificate is not valid.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "not_before",
				Description: "The time before which the certificate is not valid.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "certificate_transparency_logging_preference",
				Description: "You can opt out of certificate transparency logging by specifying the DISABLED option",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Options.CertificateTransparencyLoggingPreference"),
			},
			{
				Name:        "renewal_eligibility",
				Description: "Specifies whether the certificate is eligible for renewal",
				Type:        schema.TypeString,
			},
			{
				Name:        "renewal_summary_domain_validation_options",
				Description: "Contains information about the validation of each domain name in the certificate, as it pertains to ACM's managed renewal.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("RenewalSummary.DomainValidationOptions"),
			},
			{
				Name:        "renewal_summary_status",
				Description: "The status of ACM's managed renewal (https://docs.aws.amazon.com/acm/latest/userguide/acm-renewal.html) of the certificate. ",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RenewalSummary.RenewalStatus"),
			},
			{
				Name:        "renewal_summary_updated_at",
				Description: "The time at which the renewal summary was last updated. ",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("RenewalSummary.UpdatedAt"),
			},
			{
				Name:        "renewal_summary_failure_reason",
				Description: "The reason that a renewal request was unsuccessful.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("RenewalSummary.RenewalStatusReason"),
			},
			{
				Name:        "revocation_reason",
				Description: "The reason the certificate was revoked",
				Type:        schema.TypeString,
			},
			{
				Name:        "revoked_at",
				Description: "The time at which the certificate was revoked",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "serial",
				Description: "The serial number of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "signature_algorithm",
				Description: "The algorithm that was used to sign the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subject",
				Description: "The name of the entity that is associated with the public key contained in the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subject_alternative_names",
				Description: "One or more domain names (subject alternative names) included in the certificate",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "type",
				Description: "The source of the certificate",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tags that have been applied to the ACM certificate.",
				Type:        schema.TypeJSON,
				Resolver:    resolveACMCertificateTags,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchAcmCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ACM
	optsFn := func(o *acm.Options) {
		o.Region = cl.Region
	}
	var input acm.ListCertificatesInput
	for {
		output, err := svc.ListCertificates(ctx, &input, optsFn)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.CertificateSummaryList {
			do, err := svc.DescribeCertificate(ctx, &acm.DescribeCertificateInput{CertificateArn: item.CertificateArn}, optsFn)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- do.Certificate
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func resolveACMCertificateTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cert := resource.Item.(*types.CertificateDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().ACM
	out, err := svc.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
