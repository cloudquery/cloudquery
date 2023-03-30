package codepipeline

import (
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Pipelines() *schema.Table {
	tableName := "aws_codepipeline_pipelines"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html`,
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codepipeline"),
		Transform:           client.TransformWithStruct(&codepipeline.GetPipelineOutput{}),
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
