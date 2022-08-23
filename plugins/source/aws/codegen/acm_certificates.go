// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

func AcmCertificates() *schema.Table {
	return &schema.Table{
		Name:      "aws_acm_certificates",
		Resolver:  fetchAcmCertificates,
		Multiplex: client.ServiceAccountRegionMultiplexer("acm"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
			},
			{
				Name: "certificate_authority_arn",
				Type: schema.TypeString,
			},
			{
				Name: "created_at",
				Type: schema.TypeJSON,
			},
			{
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "domain_validation_options",
				Type: schema.TypeJSON,
			},
			{
				Name: "extended_key_usages",
				Type: schema.TypeJSON,
			},
			{
				Name: "failure_reason",
				Type: schema.TypeString,
			},
			{
				Name: "imported_at",
				Type: schema.TypeJSON,
			},
			{
				Name: "in_use_by",
				Type: schema.TypeStringArray,
			},
			{
				Name: "issued_at",
				Type: schema.TypeJSON,
			},
			{
				Name: "issuer",
				Type: schema.TypeString,
			},
			{
				Name: "key_algorithm",
				Type: schema.TypeString,
			},
			{
				Name:     "key_usages",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("KeyUsages.Name"),
			},
			{
				Name: "not_after",
				Type: schema.TypeJSON,
			},
			{
				Name: "not_before",
				Type: schema.TypeJSON,
			},
			{
				Name: "options",
				Type: schema.TypeJSON,
			},
			{
				Name: "renewal_eligibility",
				Type: schema.TypeString,
			},
			{
				Name: "renewal_summary",
				Type: schema.TypeJSON,
			},
			{
				Name: "revocation_reason",
				Type: schema.TypeString,
			},
			{
				Name: "revoked_at",
				Type: schema.TypeJSON,
			},
			{
				Name: "serial",
				Type: schema.TypeString,
			},
			{
				Name: "signature_algorithm",
				Type: schema.TypeString,
			},
			{
				Name: "status",
				Type: schema.TypeString,
			},
			{
				Name: "subject",
				Type: schema.TypeString,
			},
			{
				Name: "subject_alternative_names",
				Type: schema.TypeStringArray,
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
		},
	}
}

func fetchAcmCertificates(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().ACM

	var input acm.ListCertificatesInput
	paginator := acm.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.CertificateSummaryList {
			do, err := svc.DescribeCertificate(ctx, &acm.DescribeCertificateInput{
				CertificateArn: item.CertificateArn,
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- do.Certificate
		}
	}
	return nil
}

func resolveAcmCertificatesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cert := resource.Item.(*types.CertificateDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().ACM
	out, err := svc.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{
		CertificateArn: cert.CertificateArn,
	})
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
