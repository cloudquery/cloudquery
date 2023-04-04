package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Apps() *schema.Table {
	tableName := "aws_resiliencehub_apps"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_App.html`,
		Resolver:            fetchApps,
		PreResourceResolver: describeApp,
		Transform:           transformers.TransformWithStruct(&types.App{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns:             []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("AppArn")},
		Relations:           []*schema.Table{appAssesments(), appVersions()},
	}
}

func fetchApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Resiliencehub
	p := resiliencehub.NewListAppsPaginator(svc, &resiliencehub.ListAppsInput{})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.AppSummaries
	}
	return nil
}

func describeApp(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Resiliencehub
	out, err := svc.DescribeApp(ctx,
		&resiliencehub.DescribeAppInput{AppArn: resource.Item.(types.AppSummary).AppArn},
	)
	if err != nil {
		return err
	}
	resource.SetItem(out.App)
	return nil
}
