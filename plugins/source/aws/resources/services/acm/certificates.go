package acm

import (
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Certificates() *schema.Table {
	tableName := "aws_acm_certificates"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html`,
		Resolver:            fetchAcmCertificates,
		PreResourceResolver: getCertificate,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "acm"),
		Transform:           transformers.TransformWithStruct(&types.CertificateDetail{}, transformers.WithSkipFields("CertificateArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
