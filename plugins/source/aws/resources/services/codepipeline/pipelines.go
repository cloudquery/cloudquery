// Code generated by codegen; DO NOT EDIT.

package codepipeline

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Pipelines() *schema.Table {
	return &schema.Table{
		Name:                "aws_codepipeline_pipelines",
		Resolver:            fetchCodepipelinePipelines,
		PreResourceResolver: getPipeline,
		Multiplex:           client.ServiceAccountRegionMultiplexer("codepipeline"),
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
			{
				Name:     "metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metadata"),
			},
			{
				Name:     "pipeline",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Pipeline"),
			},
		},
	}
}
