package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CaCertificates() *schema.Table {
	tableName := "aws_iot_ca_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_CACertificateDescription.html`,
		Resolver:    fetchIotCaCertificates,
		Transform:   transformers.TransformWithStruct(&types.CACertificateDescription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "certificates",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotCaCertificateCertificates,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchIotCaCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	input := iot.ListCACertificatesInput{
		PageSize: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().Iot
	for {
		response, err := svc.ListCACertificates(ctx, &input)
		if err != nil {
			return err
		}
		for _, ca := range response.Certificates {
			cert, err := svc.DescribeCACertificate(ctx, &iot.DescribeCACertificateInput{
				CertificateId: ca.CertificateId,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- cert.CertificateDescription
		}
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
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
	for {
		response, err := svc.ListCertificatesByCA(ctx, &input)
		if err != nil {
			return err
		}

		for _, ct := range response.Certificates {
			certs = append(certs, *ct.CertificateId)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, certs)
}
