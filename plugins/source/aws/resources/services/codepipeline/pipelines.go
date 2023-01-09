package codepipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:                "aws_codepipeline_pipelines",
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer("codepipeline"),
		Transform:           transformers.TransformWithStruct(&codepipeline.GetPipelineOutput{}),
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
				Resolver: resolvePipelineArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolvePipelineTags,
			},
		},
	}
}
