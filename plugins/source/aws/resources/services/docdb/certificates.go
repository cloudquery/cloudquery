// Code generated by codegen; DO NOT EDIT.

package docdb

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_docdb_certificates",
		Description: "https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html",
		Resolver:    fetchDocdbCertificates,
		Multiplex:   client.ServiceAccountRegionMultiplexer("docdb"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "certificate_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateIdentifier"),
			},
			{
				Name:     "certificate_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CertificateType"),
			},
			{
				Name:     "thumbprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Thumbprint"),
			},
			{
				Name:     "valid_from",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ValidFrom"),
			},
			{
				Name:     "valid_till",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ValidTill"),
			},
		},
	}
}
