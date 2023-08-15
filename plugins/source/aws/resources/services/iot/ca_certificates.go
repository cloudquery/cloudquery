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

func CaCertificates() *schema.Table {
	tableName := "aws_iot_ca_certificates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_CACertificateDescription.html`,
		Resolver:            fetchIotCaCertificates,
		PreResourceResolver: getCaCertificate,
		Transform:           transformers.TransformWithStruct(&types.CACertificateDescription{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "certificates",
				Type:     arrow.ListOf(arrow.BinaryTypes.String),
				Resolver: ResolveIotCaCertificateCertificates,
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

func fetchIotCaCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListCACertificatesInput{
		PageSize: aws.Int32(250),
	}
	cl := meta.(*client.Client)

	svc := cl.Services().Iot
	paginator := iot.NewListCACertificatesPaginator(svc, &input)
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

func getCaCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot

	output, err := svc.DescribeCACertificate(ctx, &iot.DescribeCACertificateInput{
		CertificateId: resource.Item.(types.CACertificate).CertificateId,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output.CertificateDescription
	return nil
}

func ResolveIotCaCertificateCertificates(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*types.CACertificateDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListCertificatesByCAInput{
		CaCertificateId: i.CertificateId,
		PageSize:        aws.Int32(250),
	}
	var certs []string
	paginator := iot.NewListCertificatesByCAPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, ct := range page.Certificates {
			certs = append(certs, *ct.CertificateId)
		}
	}
	return resource.Set(c.Name, certs)
}
