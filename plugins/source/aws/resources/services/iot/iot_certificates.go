package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func IotCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_certificates",
		Description: "Describes a certificate.",
		Resolver:    fetchIotCertificates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
				Name:        "policies",
				Description: "Policies of the certificate",
				Type:        schema.TypeStringArray,
				Resolver:    ResolveIotCertificatePolicies,
			},
			{
				Name:          "ca_certificate_id",
				Description:   "The certificate ID of the CA certificate used to sign this certificate.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:            "arn",
				Description:     "The ARN of the certificate.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("CertificateArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "id",
				Description: "The ID of the certificate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateId"),
			},
			{
				Name:        "mode",
				Description: "The mode of the certificate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateMode"),
			},
			{
				Name:        "pem",
				Description: "The certificate data, in PEM format.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificatePem"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time the certificate was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "customer_version",
				Description: "The customer version of the certificate.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "generation_id",
				Description: "The generation ID of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_modified_date",
				Description: "The date and time the certificate was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "owned_by",
				Description: "The ID of the Amazon Web Services account that owns the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:          "previous_owned_by",
				Description:   "The ID of the Amazon Web Services account of the previous owner of the certificate.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "The status of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:          "transfer_data",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("TransferData"),
				IgnoreInTests: true,
			},
			{
				Name:     "validity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Validity"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchIotCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListCertificatesInput{
		PageSize: aws.Int32(250),
	}

	for {
		response, err := svc.ListCertificates(ctx, &input)
		if err != nil {
			return err
		}

		for _, ct := range response.Certificates {
			cert, err := svc.DescribeCertificate(ctx, &iot.DescribeCertificateInput{
				CertificateId: ct.CertificateId,
			}, func(options *iot.Options) {
				options.Region = cl.Region
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
func ResolveIotCertificatePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*types.CertificateDescription)
	cl := meta.(*client.Client)
	svc := cl.Services().IOT
	input := iot.ListAttachedPoliciesInput{
		Target:   i.CertificateArn,
		PageSize: aws.Int32(250),
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
