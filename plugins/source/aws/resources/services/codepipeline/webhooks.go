package codepipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Webhooks() *schema.Table {
	return &schema.Table{
		Name:        "aws_codepipeline_webhooks",
		Description: `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html`,
		Resolver:    fetchCodepipelineWebhooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("codepipeline"),
		Transform: transformers.TransformWithStruct(&types.ListWebhookItem{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
