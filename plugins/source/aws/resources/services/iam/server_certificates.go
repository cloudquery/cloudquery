// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ServerCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_server_certificates",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html`,
		Resolver:    fetchIamServerCertificates,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCertificateId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Path"),
			},
			{
				Name:     "server_certificate_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServerCertificateName"),
			},
			{
				Name:     "expiration",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Expiration"),
			},
			{
				Name:     "upload_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UploadDate"),
			},
		},
	}
}
