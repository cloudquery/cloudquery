package iam

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func signingCertificates() *schema.Table {
	tableName := "aws_iam_signing_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html`,
		Resolver:    fetchUserSigningCertificates,
		Transform:   transformers.TransformWithStruct(&types.SigningCertificate{}, transformers.WithPrimaryKeys("CertificateId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:       "user_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
		},
	}
}

func fetchUserSigningCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := &iam.ListSigningCertificatesInput{UserName: parent.Item.(*types.User).UserName}
	cl := meta.(*client.Client)
	svc := cl.Services().Iam
	paginator := iam.NewListSigningCertificatesPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Certificates
	}
	return nil
}
