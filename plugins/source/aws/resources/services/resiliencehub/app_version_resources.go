package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func appVersionResources() *schema.Table {
	tableName := "aws_resiliencehub_app_version_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_PhysicalResource.html`,
		Resolver:    fetchAppVersionResources,
		Transform:   transformers.TransformWithStruct(&types.PhysicalResource{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, appVersion,
			{
				Name:            "physical_resource_identifier",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("PhysicalResourceId.Identifier"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}

func fetchAppVersionResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Resiliencehub
	p := resiliencehub.NewListAppVersionResourcesPaginator(svc, &resiliencehub.ListAppVersionResourcesInput{
		AppArn:     parent.Parent.Item.(*types.App).AppArn,
		AppVersion: parent.Item.(types.AppVersionSummary).AppVersion,
		MaxResults: aws.Int32(100),
	})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx)
		if err != nil {
			return err
		}

		res <- out.PhysicalResources
	}
	return nil
}
