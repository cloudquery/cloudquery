package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchCodepipelineWebhooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Codepipeline
	paginator := codepipeline.NewListWebhooksPaginator(svc, &codepipeline.ListWebhooksInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Webhooks
	}
	return nil
}
