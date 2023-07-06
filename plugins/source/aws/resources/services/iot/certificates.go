package iot

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Certificates() *schema.Table {
	tableName := "aws_iot_certificates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html`,
		Resolver:            fetchIotCertificates,
		PreResourceResolver: getCertificate,
		Transform:           transformers.TransformWithStruct(&types.CertificateDescription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "policies",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: ResolveIotCertificatePolicies,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("CertificateArn"),
				PrimaryKey: true,
			},
		},
	}
}
func fetchIotCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListCertificatesInput{
		PageSize: aws.Int32(250),
	}
	paginator := iot.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Certificates
	}
	return nil
}

func getCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cert := resource.Item.(types.Certificate)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	certDescription, err := svc.DescribeCertificate(ctx, &iot.DescribeCertificateInput{
		CertificateId: cert.CertificateId,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = certDescription.CertificateDescription
	return nil
}

func ResolveIotCertificatePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListAttachedPoliciesInput{
		Target:   resource.Item.(*types.CertificateDescription).CertificateArn,
		PageSize: aws.Int32(250),
	}
	paginator := iot.NewListAttachedPoliciesPaginator(svc, &input)
	var policies []string
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}
	}
	return resource.Set(c.Name, policies)
}
