// Code generated by codegen; DO NOT EDIT.

package elastictranscoder

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PipelineJobs() *schema.Table {
	return &schema.Table{
		Name:        "aws_elastictranscoder_pipeline_jobs",
		Description: `https://docs.aws.amazon.com/elastictranscoder/latest/developerguide/list-jobs-by-pipeline.html`,
		Resolver:    fetchElastictranscoderPipelineJobs,
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "input",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Input"),
			},
			{
				Name:     "inputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Inputs"),
			},
			{
				Name:     "output",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Output"),
			},
			{
				Name:     "output_key_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutputKeyPrefix"),
			},
			{
				Name:     "outputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Outputs"),
			},
			{
				Name:     "pipeline_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PipelineId"),
			},
			{
				Name:     "playlists",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Playlists"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "timing",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Timing"),
			},
			{
				Name:     "user_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserMetadata"),
			},
		},
	}
}
