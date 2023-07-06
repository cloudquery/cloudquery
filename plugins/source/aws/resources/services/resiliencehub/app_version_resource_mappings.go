package resiliencehub

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func appVersionResourceMappings() *schema.Table {
	tableName := "aws_resiliencehub_app_version_resource_mappings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResourceMapping.html`,
		Resolver:    fetchAppVersionResourceMappings,
		Transform:   transformers.TransformWithStruct(&types.ResourceMapping{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "resiliencehub"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, appVersion,
			{
				Name:       "physical_resource_identifier",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("PhysicalResourceId.Identifier"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAppVersionResourceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Resiliencehub
	p := resiliencehub.NewListAppVersionResourceMappingsPaginator(svc, &resiliencehub.ListAppVersionResourceMappingsInput{
		AppArn:     parent.Parent.Item.(*types.App).AppArn,
		AppVersion: parent.Item.(types.AppVersionSummary).AppVersion,
	})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- out.ResourceMappings
	}
	return nil
}
