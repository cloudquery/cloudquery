package acmpca

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func CertificateAuthorities() *schema.Table {
	tableName := "aws_acmpca_certificate_authorities"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/privateca/latest/APIReference/API_CertificateAuthority.html`,
		Resolver:    fetchAcmpcaCertificateAuthorities,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "acm-pca"),
		Transform:   transformers.TransformWithStruct(&types.CertificateAuthority{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Arn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     cqtypes.ExtensionTypes.JSON,
				Resolver: resolveCertificateAuthorityTags,
			},
		},
	}
}

func fetchAcmpcaCertificateAuthorities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acmpca
	paginator := acmpca.NewListCertificateAuthoritiesPaginator(svc, nil)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *acmpca.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.CertificateAuthorities
	}
	return nil
}

func resolveCertificateAuthorityTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	certAuthority := resource.Item.(types.CertificateAuthority)
	cl := meta.(*client.Client)
	svc := cl.Services().Acmpca
	out, err := svc.ListTags(ctx,
		&acmpca.ListTagsInput{CertificateAuthorityArn: certAuthority.Arn},
		func(o *acmpca.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.Tags))
}
