package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SigningCertificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_signing_certificates",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html`,
		Resolver:    fetchUserSigningCertificates,
		Transform:   transformers.TransformWithStruct(&types.SigningCertificate{}, transformers.WithPrimaryKeys("CertificateId")),
		Multiplex:   client.ServiceAccountRegionMultiplexer("iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "user_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
