package codepipeline

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Webhooks() *schema.Table {
	tableName := "aws_codepipeline_webhooks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html`,
		Resolver:    fetchCodepipelineWebhooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codepipeline"),
		Transform:   transformers.TransformWithStruct(&types.ListWebhookItem{}),
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
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchCodepipelineWebhooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Codepipeline
	paginator := codepipeline.NewListWebhooksPaginator(svc, &codepipeline.ListWebhooksInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codepipeline.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Webhooks
	}
	return nil
}
