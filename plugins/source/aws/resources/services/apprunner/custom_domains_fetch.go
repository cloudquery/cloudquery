package apprunner

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchApprunnerCustomDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	paginator := apprunner.NewDescribeCustomDomainsPaginator(meta.(*client.Client).Services().Apprunner,
		&apprunner.DescribeCustomDomainsInput{ServiceArn: parent.Item.(*types.Service).ServiceArn})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.CustomDomains
	}
	return nil
}
