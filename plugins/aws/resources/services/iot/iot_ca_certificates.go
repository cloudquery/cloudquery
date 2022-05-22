package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func IotCaCertificates() *schema.Table {
	return &schema.Table{
		Name:          "aws_iot_ca_certificates",
		Description:   "Describes a CA certificate.",
		Resolver:      fetchIotCaCertificates,
		Multiplex:     client.ServiceAccountRegionMultiplexer("iot"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "certificates",
				Description: "Certificates of the ca certificate",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveIotCaCertificateCertificates,
			},
			{
				Name:        "auto_registration_status",
				Description: "Whether the CA certificate configured for auto registration of device certificates",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The CA certificate ARN.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateArn"),
			},
			{
				Name:        "id",
				Description: "The CA certificate ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateId"),
			},
			{
				Name:        "pem",
				Description: "The CA certificate data, in PEM format.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificatePem"),
			},
			{
				Name:        "creation_date",
				Description: "The date the CA certificate was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "customer_version",
				Description: "The customer version of the CA certificate.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "generation_id",
				Description: "The generation ID of the CA certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_date",
				Description: "The date the CA certificate was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "owned_by",
				Description: "The owner of the CA certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of a CA certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "validity_not_after",
				Description: "The certificate is not valid after this date.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Validity.NotAfter"),
			},
			{
				Name:        "validity_not_before",
				Description: "The certificate is not valid before this date.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Validity.NotBefore"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotCaCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	input := iot.ListCACertificatesInput{
		PageSize: aws.Int32(250),
	}
	c := meta.(*client.Client)

	svc := c.Services().IOT
	for {
		response, err := svc.ListCACertificates(ctx, &input, func(options *iot.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		for _, ca := range response.Certificates {
			cert, err := svc.DescribeCACertificate(ctx, &iot.DescribeCACertificateInput{
				CertificateId: ca.CertificateId,
			}, func(options *iot.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return diag.WrapError(err)
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
	svc := cl.Services().IOT
	input := iot.ListCertificatesByCAInput{
		CaCertificateId: i.CertificateId,
		PageSize:        aws.Int32(250),
	}

	var certs []string
	for {
		response, err := svc.ListCertificatesByCA(ctx, &input, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return diag.WrapError(err)
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
