package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func signingCertificates() *schema.Table {
	tableName := "aws_iam_signing_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html`,
		Resolver:    fetchUserSigningCertificates,
		Transform:   transformers.TransformWithStruct(&types.SigningCertificate{}, transformers.WithPrimaryKeys("CertificateId")),
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

func fetchUserSigningCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := &iam.ListSigningCertificatesInput{UserName: parent.Item.(*types.User).UserName}
	svc := meta.(*client.Client).Services().Iam
	paginator := iam.NewListSigningCertificatesPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Certificates
	}
	return nil
}
