// Code generated by codegen; DO NOT EDIT.

package iot

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CaCertificates() *schema.Table {
	return &schema.Table{
		Name:      "aws_iot_ca_certificates",
		Resolver:  fetchIotCaCertificates,
		Multiplex: client.ServiceAccountRegionMultiplexer("iot"),
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
			{
				Name:     "auto_registration_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoRegistrationStatus"),
			},
			{
				Name:     "certificate_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateId"),
			},
			{
				Name:     "certificate_pem",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificatePem"),
			},
			{
				Name:     "creation_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationDate"),
			},
			{
				Name:     "customer_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CustomerVersion"),
			},
			{
				Name:     "generation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GenerationId"),
			},
			{
				Name:     "last_modified_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedDate"),
			},
			{
				Name:     "owned_by",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnedBy"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "validity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Validity"),
			},
		},
	}
}
