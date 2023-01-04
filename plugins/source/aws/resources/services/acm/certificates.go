package acm

import (
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:                "aws_acm_certificates",
		Description:         `https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html`,
		Resolver:            fetchAcmCertificates,
		PreResourceResolver: getCertificate,
		Multiplex:           client.ServiceAccountRegionMultiplexer("acm"),
		Transform:           transformers.TransformWithStruct(&types.CertificateDetail{}),
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCertificateTags,
			},
		},
	}
}
