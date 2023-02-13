package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CaCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_ca_certificates",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_CACertificateDescription.html`,
		Resolver:    fetchIotCaCertificates,
		Transform:   transformers.TransformWithStruct(&types.CACertificateDescription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
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
