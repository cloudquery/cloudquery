package iot

import (
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_iot_certificates",
		Description: `https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html`,
		Resolver:    fetchIotCertificates,
		Transform:   transformers.TransformWithStruct(&types.CertificateDescription{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "policies",
				Type:     schema.TypeStringArray,
				Resolver: ResolveIotCertificatePolicies,
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
