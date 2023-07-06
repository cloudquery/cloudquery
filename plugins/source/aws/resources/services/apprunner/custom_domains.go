package apprunner

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func customDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_apprunner_custom_domains",
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html`,
		Resolver:    fetchApprunnerCustomDomains,
		Transform:   transformers.TransformWithStruct(&types.CustomDomain{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "enable_www_subdomain",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("EnableWWWSubdomain"),
			},
		},
	}
}

func fetchApprunnerCustomDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	paginator := apprunner.NewDescribeCustomDomainsPaginator(meta.(*client.Client).Services().Apprunner,
		&apprunner.DescribeCustomDomainsInput{ServiceArn: parent.Item.(*types.Service).ServiceArn})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.CustomDomains
	}
	return nil
}
