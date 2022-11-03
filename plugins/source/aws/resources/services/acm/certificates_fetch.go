package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchAcmCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	var input acm.ListCertificatesInput
	paginator := acm.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.CertificateSummaryList
	}
	return nil
}

func getCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	input := acm.DescribeCertificateInput{CertificateArn: resource.Item.(types.CertificateSummary).CertificateArn}
	output, err := svc.DescribeCertificate(ctx, &input)
	if err != nil {
		return err
	}
	resource.Item = output.Certificate
	return nil
}

func resolveCertificateTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cert := resource.Item.(*types.CertificateDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	out, err := svc.ListTagsForCertificate(ctx, &acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.Tags))
}
