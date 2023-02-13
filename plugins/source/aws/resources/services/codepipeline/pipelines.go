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
		Description:         `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html`,
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer("codepipeline"),
		Transform:           transformers.TransformWithStruct(&codepipeline.GetPipelineOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
