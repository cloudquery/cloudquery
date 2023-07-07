package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func appVersions() *schema.Table {
	tableName := "aws_resiliencehub_app_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_AppVersionSummary.html`,
		Resolver:    fetchAppVersions,
		Transform:   transformers.TransformWithStruct(&types.AppVersionSummary{}, transformers.WithPrimaryKeys("AppVersion")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARNTop},
		Relations:   []*schema.Table{appVersionResources(), appVersionResourceMappings()},
	}
}

func fetchAppVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resiliencehub
	p := resiliencehub.NewListAppVersionsPaginator(svc, &resiliencehub.ListAppVersionsInput{AppArn: parent.Item.(*types.App).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.AppVersions
	}
	return nil
}
