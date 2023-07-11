package acm

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("CertificateArn"),
				PrimaryKey: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCertificateTags,
			},
		},
	}
}

func fetchAcmCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	input := acm.ListCertificatesInput{
		CertificateStatuses: types.CertificateStatus("").Values(),
		Includes: &types.Filters{
			ExtendedKeyUsage: []types.ExtendedKeyUsageName{types.ExtendedKeyUsageNameAny},
			KeyTypes:         types.KeyAlgorithm("").Values(),
			KeyUsage:         []types.KeyUsageName{types.KeyUsageNameAny},
		},
	}
	paginator := acm.NewListCertificatesPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(o *acm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.CertificateSummaryList
	}
	return nil
}

func getCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	input := acm.DescribeCertificateInput{CertificateArn: resource.Item.(types.CertificateSummary).CertificateArn}
	output, err := svc.DescribeCertificate(ctx, &input, func(o *acm.Options) { o.Region = cl.Region })
	if err != nil {
		return err
	}
	resource.Item = output.Certificate
	return nil
}

func resolveCertificateTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cert := resource.Item.(*types.CertificateDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().Acm
	out, err := svc.ListTagsForCertificate(ctx,
		&acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn},
		func(o *acm.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(out.Tags))
}
