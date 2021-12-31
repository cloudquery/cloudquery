package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func RdsCertificates() *schema.Table {
	return &schema.Table{
		Name:         "aws_rds_certificates",
		Description:  "A CA certificate for an AWS account.",
		Resolver:     fetchRdsCertificates,
		Multiplex:    client.ServiceAccountRegionMultiplexer("rds"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "thumbprint"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the certificate.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CertificateArn"),
			},
			{
				Name:        "certificate_identifier",
				Description: "The unique key that identifies a certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "certificate_type",
				Description: "The type of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "customer_override",
				Description: "Whether there is an override for the default certificate identifier.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "customer_override_valid_till",
				Description: "If there is an override for the default certificate identifier, when the override expires.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "thumbprint",
				Description: "The thumbprint of the certificate.",
				Type:        schema.TypeString,
			},
			{
				Name:        "valid_from",
				Description: "The starting date from which the certificate is valid.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "valid_till",
				Description: "The final date that the certificate continues to be valid.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRdsCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config rds.DescribeCertificatesInput
	c := meta.(*client.Client)
	svc := c.Services().RDS
	for {
		response, err := svc.DescribeCertificates(ctx, &config, func(o *rds.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.Certificates
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}
